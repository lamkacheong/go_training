dao层:
RecordNotFound = errors.New("business not found error")

err = stmtOut.QueryRow(10).Scan(&user.id,&user.name) // WHERE number = 13
if err != nil {
	if err == sql.ErrNoRows {
		return errors.wrap(RecordNotFound, "sql error")
	}
}

business层:
result, err := dao()
if errors.Is(err, RecordNotFound){
	handle_not_found()
}
