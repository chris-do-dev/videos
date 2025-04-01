conn, _ := sql.Open("libsql", "file:"+filename)

// we passed an invalid conn here because we didn't check the previous error
Exec(ctx, conn, "CREATE TABLE records")
