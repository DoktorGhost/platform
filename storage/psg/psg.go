package psg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"os"
	"path/filepath"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DBConfig struct {
	DbHost  string
	DbPort  string
	DbName  string
	DbLogin string
	DbPass  string
}

func InitStorage(conf DBConfig) (*pgxpool.Pool, error) {
	login := conf.DbLogin
	password := conf.DbPass
	host := conf.DbHost
	port := conf.DbPort
	dbname := conf.DbName

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", login, password, host, port, dbname)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к бд: %v", err)
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к бд: %v", err)
	}

	// Получаем путь к папке с миграциями
	migrationsDir, err := filepath.Abs("migrations")
	if err != nil {
		return nil, fmt.Errorf("ошибка определения пути к миграциям: %v", err)
	}

	// Проверка существования директории
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("директория миграций не существует: %v", migrationsDir)
	}

	// Применяем миграции
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД для миграций: %v", err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		return nil, fmt.Errorf("ошибка применения миграций: %v", err)
	}

	return dbpool, nil
}
