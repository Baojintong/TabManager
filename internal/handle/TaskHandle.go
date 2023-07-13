package handle

import "tabManager/internal/define"

func CreateTasktable() {
	db.ExecNoTran(define.CREATE_TASK_TABLE)
}
