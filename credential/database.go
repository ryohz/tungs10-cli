package credential

import (
	"database/sql"
	"fmt"
	constant "ryozk/tungs10-cli/constant"
	"ryozk/tungs10-cli/user"
	"strconv"
)

func (credential *Credential) RemoveFromDB() error {
	userDir, gudErr := credential.user.GetUserDir()
	if gudErr != nil {
		return gudErr
	}
	var dbPath string = fmt.Sprintf("%s/%s", userDir, constant.CREDENTIAL_DATABASE_NAME)
	db, doErr := sql.Open("sqlite3", dbPath)
	if doErr != nil {
		return doErr
	}
	defer db.Close()
	var query string = "delete from credentials where id = ?"
	statement, pErr := db.Prepare(query)
	if pErr != nil {
		return pErr
	}
	defer statement.Close()
	_, seErr := statement.Exec(credential.id)
	if seErr != nil {
		return seErr
	}
	return nil
}

func (credential *Credential) AddToDB() error {
	userDir, getUserDirErr := credential.user.GetUserDir()
	if getUserDirErr != nil {
		return getUserDirErr
	}

	var dbPath string = fmt.Sprintf("%s/%s", userDir, constant.CREDENTIAL_DATABASE_NAME)
	db, doErr := sql.Open("sqlite3", dbPath)
	if doErr != nil {
		return doErr
	}
	defer db.Close()

	var sqlQry string = `
			create table credentials (id integer not null primary key autoincrement,serviceName string not null, userName string, email string, password string not null);
			delete from credentials;
		`

	_, execErr := db.Exec(sqlQry)
	if execErr != nil {
		fmt.Println(execErr)
	}

	context, bErr := db.Begin()
	if bErr != nil {
		return bErr
	}

	statement, prepareErr := context.Prepare("insert into credentials(serviceName,userName,email,password) values(?,?,?,?)")
	if prepareErr != nil {
		return prepareErr
	}
	_, eiErr := statement.Exec(credential.GetEncryptedServiceName(), credential.GetEncryptedUserName(), credential.GetEncryptedEmail(), credential.GetEncryptedPassword())
	if eiErr != nil {
		fmt.Println(eiErr)
	}
	defer statement.Close()

	var commitErr error = context.Commit()
	if commitErr != nil {
		return commitErr
	}
	return nil
}

func (credential *Credential) ReadAndSetEncryptedCredentialInformations() error {
	userDir, gudErr := credential.user.GetUserDir()
	if gudErr != nil {
		return gudErr
	}
	var dbPath string = fmt.Sprintf("%s/%s", userDir, constant.CREDENTIAL_DATABASE_NAME)
	db, doErr := sql.Open("sqlite3", dbPath)
	if doErr != nil {
		return doErr
	}
	defer db.Close()

	if credential.id != 0 {
		statement, dpErr := db.Prepare("select * from credentials where id = ?")
		if dpErr != nil {
			fmt.Println(dpErr)
		}
		defer statement.Close()
		var row *sql.Row = statement.QueryRow(strconv.Itoa(credential.GetId()))
		var id int
		var encryptedServiceName string
		var encryptedUserName string
		var encryptedEmail string
		var encryptedPassword string
		gcErr := row.Scan(&id, &encryptedServiceName, &encryptedUserName, &encryptedEmail, &encryptedPassword)
		if gcErr != nil {
			return gcErr
		}
		credential.SetId(id)
		credential.SetEncryptedServiceName(encryptedServiceName)
		credential.SetEncryptedUserName(encryptedUserName)
		credential.SetEncryptedEmail(encryptedEmail)
		credential.SetEncryptedPassword(encryptedPassword)

		return nil
	}
	//filtering handling
	return nil
}

func GetLastId(user user.User) (int, error) {
	userDir, gudErr := user.GetUserDir()
	if gudErr != nil {
		return 0, gudErr
	}
	var dbPath string = fmt.Sprintf("%s/%s", userDir, constant.CREDENTIAL_DATABASE_NAME)
	db, doErr := sql.Open("sqlite3", dbPath)
	if doErr != nil {
		return 0, doErr
	}
	defer db.Close()
	row, pErr := db.Prepare("select max(id) from credentials")
	if pErr != nil {
		return 0, pErr
	}
	var id int
	var sErr error = row.QueryRow().Scan(&id)
	if sErr != nil {
		return 0, sErr
	}
	return id, nil
}

func IsCredentialExists(user user.User) (bool, error) {
	userDir, gudErr := user.GetUserDir()
	if gudErr != nil {
		return false, gudErr
	}
	var dbPath string = fmt.Sprintf("%s/%s", userDir, constant.CREDENTIAL_DATABASE_NAME)
	db, doErr := sql.Open("sqlite3", dbPath)
	if doErr != nil {
		return false, doErr
	}
	defer db.Close()
	statement, pErr := db.Prepare("select count(*) from credentials ")
	if pErr != nil {
		return false, pErr
	}
	var rowSum int
	statement.QueryRow().Scan(&rowSum)
	if rowSum > 0 {
		return true, nil
	}
	return false, nil
}
