all: 
	go install gs_main/game_server.go 
	go install gate_main/gate_server.go
p:
	bash make_proto.sh
