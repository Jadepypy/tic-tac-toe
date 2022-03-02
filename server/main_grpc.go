package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"tic_tac_toe/server/model"
	pb "tic_tac_toe/tic_tac_toe_proto"
)

type server struct {
	pb.UnimplementedPlayerServer
}

func (s *server) GetRecords(ctx context.Context, empty *emptypb.Empty) (*pb.Records, error) {
	stateIDArr, userIDArr := gameController.GetRecotds()
	records := make([]*pb.Record, 0)
	for i, _ := range stateIDArr {
		stateID := int32(stateIDArr[i])
		userID := int32(userIDArr[i])
		record := pb.Record{State: &stateID, Player: &userID}
		records = append(records, &record)
		//fmt.Println(record.GetPlayer(), record.GetState())
	}
	return &pb.Records{RecordList: records}, nil
}

func (s *server) Play(ctx context.Context, playerMove *pb.PlayerMove) (*emptypb.Empty, error) {
	log.Printf("Received")
	xIndex := int(playerMove.GetXIndex())
	yIndex := int(playerMove.GetYIndex())
	fmt.Println(xIndex, yIndex)
	gameController.Play(xIndex, yIndex)
	return new(emptypb.Empty), nil
}

func (s *server) GetBoard(ctx context.Context, empty *emptypb.Empty) (*pb.GameState, error) {
	log.Printf("Received")
	boardArr, state, turn := gameController.GetGame()
	boardFlat := make([]string, 9)
	for i := range boardArr {
		for j := range boardArr[0] {
			boardFlat[i+j*3] = boardArr[i][j]
		}
	}
	fmt.Println(boardFlat, state, turn)
	gameState := &pb.GameState{Board: boardFlat, State: int32(state), Turn: int32(turn)}
	return gameState, nil
}

func MainGRPC() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("fail to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPlayerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	model.InitDB()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
