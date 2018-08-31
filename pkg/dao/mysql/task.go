package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/leopoldxx/go-utils/trace"
	"reboot/pkg/dao/mysql/types"
)

func updateTableFieldByID(ctx context.Context, db *sqlx.DB, tx *sqlx.Tx, table string, id int64, fields map[types.Field]types.Value) error {
	var (
		sql = `
								    		    		UPDATE %s 
								    		    		SET %s
								    		    		WHERE id=? ;`
	)
	tracer := trace.GetTraceFromContext(ctx)
	//if len(fields) == 0 {
	//	tracer.Warnf("invalid update fields")
	//	return errors.New("invalid update fields")
	//}
	var (
		dest        []string
		fieldsValue []interface{}
		err         error
	)
	for k, v := range fields {
		dest = append(dest, fmt.Sprintf("%s=?", string(k)))
		fieldsValue = append(fieldsValue, v)
	}
	// append the last id field in where clause
	fieldsValue = append(fieldsValue, id)
	sql = fmt.Sprintf(sql, table, strings.Join(dest, ","))

	if db != nil {
		_, err = db.Exec(sql, fieldsValue...)
	} else {
		_, err = tx.Exec(sql, fieldsValue...)
	}
	if err != nil {
		tracer.Errorf("failed to update %s: %s", table, err)
		return err
	}

	tracer.Infof("update table %s #%d successfully", table, id)
	return nil
}

// update task status
func (m *mysql) UpdateTask(ctx context.Context, task *types.Task) error {
	return updateTableFieldByID(ctx, m.db, nil, "task", task.ID, map[types.Field]types.Value{
		types.FieldStatus:   task.Status,
		types.FieldIsClosed: task.IsClosed,
		types.FieldIsPaused: task.IsPaused,
	})
}

func (m *mysql) CreateTask(ctx context.Context, task *types.Task) (int64, error) {
	const (
		sqlTpl = `
    		INSERT INTO task (
    		namespace,
    		resource,
    		task_type,
    		spec,
    		status,
    		op_user,
    		create_time)
	        VALUES (:namespace,
		:resource,
		:task_type,
	        :spec,
		:status,
		:op_user,
		NOW());`
	)

	tracer := trace.GetTraceFromContext(ctx)

	var (
		res sql.Result
		err error
	)
	tracer.Info(sqlTpl)
	tracer.Info(*task)

	if m.db != nil {
		res, err = m.db.NamedExec(sqlTpl, *task)
		if err != nil {
			tracer.Errorf("failed to insert task: %s", err)
			return 0, err
		}
	}

	tracer.Info("insert task successfully")
	lastID, err := res.LastInsertId()
	if err != nil {
		tracer.Errorf("failed to get lastid of the task: %s", err)
		return 0, err
	}
	return lastID, err
}

func (m *mysql) ListOpenTasks(ctx context.Context) ([]types.Task, error) {
	tasks, err := getTaskByField(ctx, m.db, map[types.Field]types.Value{
		types.FieldIsClosed: false,
		types.FieldIsPaused: false,
	})
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (m *mysql) ListTask(ctx context.Context) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("CreateTask")
}

func (m *mysql) GetTask(ctx context.Context) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("CreateTask")
}

func (m *mysql) DeleteTask(ctx context.Context) {
	tracer := trace.GetTraceFromContext(ctx)
	tracer.Info("CreateTask")
}

func (m *mysql) GetOpenTaskByTaskID(ctx context.Context, id int64) (*types.Task, error) {
	tasks, err := getTaskByField(ctx, m.db, map[types.Field]types.Value{
		types.FieldID:       id,
		types.FieldIsClosed: false,
	})
	if err != nil {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, errors.New("task not found")
	}

	task := &tasks[0]
	if !task.IsPaused || (task.IsPaused && task.IsSkipPaused) {
		return task, nil
	}

	return nil, errors.New("task not found")
}

func getTaskByField(ctx context.Context, db *sqlx.DB,
	fields map[types.Field]types.Value) ([]types.Task, error) {
	const (
		sql = ` SELECT id,namespace, resource, task_type, spec, status, is_canceled, is_paused, is_skip_paused,
																																																																																																																																															    	    	    			is_urgent_skipped, is_closed, is_closed_manually, op_user, create_time, last_update_time
																																																																																																																																															    	    	    			FROM task
																																																																																																																																															    	    	    			WHERE %s
																																																																																																																																															    	    	    			ORDER BY create_time DESC ;`
		//	AND %s = ?
		//LIMIT 1;`
	)
	tasks := []types.Task{}
	err := getTableRowsByField(ctx, db, sql, &tasks, fields)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func getTableRowsByField(ctx context.Context, db *sqlx.DB, sql string,
	result interface{}, fields map[types.Field]types.Value) error {

	tracer := trace.GetTraceFromContext(ctx)
	var fieldsValue []interface{}
	if len(fields) > 0 {
		sql, fieldsValue = formatQuerySQL(sql, fields)
	}
	//tracer.Info(sql, fieldsValue)

	err := db.Select(result, sql, fieldsValue...)
	if err != nil {
		tracer.Errorf("failed to get data: %v", err)
		return err
	}
	return nil
}

func formatQuerySQL(sql string, fields map[types.Field]types.Value) (string, []interface{}) {
	var (
		dest        []string
		fieldsValue []interface{}
	)
	for k, v := range fields {
		value := reflect.ValueOf(v)
		if value.Kind() == reflect.Slice {
			slen := value.Len()
			if slen > 0 {
				placeholder := []string{}
				valueholder := []interface{}{}
				for i := 0; i < slen; i++ {
					placeholder = append(placeholder, "?")
					valueholder = append(valueholder,
						value.Index(i).Interface())
				}
				dest = append(dest, fmt.Sprintf("%s IN (%s)",
					string(k), strings.Join(placeholder, ",")))
				fieldsValue = append(fieldsValue, valueholder...)
			}
		} else {
			if fv, ok := v.(types.Field); ok {
				dest = append(dest, fmt.Sprintf("%s=%s", string(k), string(fv)))
			} else {
				dest = append(dest, fmt.Sprintf("%s=?", string(k)))
				fieldsValue = append(fieldsValue, v)
			}
		}
	}

	if len(dest) == 0 {
		return sql, fieldsValue
	}

	sql = fmt.Sprintf(sql, strings.Join(dest, " AND "))
	return sql, fieldsValue
}
