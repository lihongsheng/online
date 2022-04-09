package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"online/common/log"
	"online/server/api"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

var (
	sigExitOnce = new(sync.Once)
)

func init() {
	// 这段代码有屌用
	go sigExitOnce.Do(func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		defer signal.Stop(c)

		for {
			select {
			case <-c:
				fmt.Printf("exit by signal [SIGTERM/SIGINT/SIGKILL]")
				os.Exit(1)
			}
		}
	})
}

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{}
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port",
			Value: 80,
		},
		cli.StringFlag{
			Name: "fe",
		},
	}

	app.Before = func(context *cli.Context) error {
		level, err := log.ParseLevel("info")
		if err != nil {
			return err
		}
		log.SetLevel(level)
		log.SetOutput(os.Stdout)
		return nil
	}

	app.Action = func(c *cli.Context) error {
		println("------------Yakit Online Banner-----------")
		// 生成 链接 Postgres 的信息
		pg := api.GeneratePostgresParams("127.0.0.1", 5432, "root", "password")
		log.Info("start to run server")
		// 启动服务
		err := api.StartServer(
			pg, // Postgres 数据库链接信息
			c.Int("port"),  // web port
			c.String("fe"), // frontend dir
		)
		if err != nil {
			log.Errorf("start yaklang.online service failed: %v", err)
			return err
		}
		ctx := context.Background()
		select {
		case <-ctx.Done():
		}

		return errors.New("server finished")
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("command: [%v] failed: %v\n", strings.Join(os.Args, " "), err)
		return
	}
}
