#ifndef _RPCWORKER_HPP
#define _RPCWORKER_HPP

#include <grpc++/grpc++.h>
#include "remotexec.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::ServerReader;
using grpc::Channel;
using grpc::ClientContext;
using grpc::ClientWriter;
using grpc::Status;
using com::gabecloud::hypercube::scripts::ExecuteRequest;
using com::gabecloud::hypercube::scripts::ExecuteReply;
using com::gabecloud::hypercube::scripts::FileChunk;
using com::gabecloud::hypercube::scripts::TransmitAck;
using com::gabecloud::hypercube::scripts::RemoteExec;

#include "netlayer.hpp"
#include "exception.hpp"
#include "threadpool.hpp"
using namespace std;

class RemoteExecServiceImpl final : public RemoteExec::Service
{
    private:
        Status Execute(ServerContext* context, const ExecuteRequest* request, ExecuteReply* reply) override;
        Status Transmit(ServerContext* context, ServerReader<FileChunk>* reader, TransmitAck* ack) override;
        int exec_cmd(int id, char *cmd);
        NetLayer & sciNet;
        bool running;
    public:
        RemoteExecServiceImpl(NetLayer & sci);
        bool getState() { return running; }
};

class FrontBack {
    public:
        FrontBack(shared_ptr<Channel> channel)
            : stub_(RemoteExec::NewStub(channel)) {}
        string Execute(int be_id, int msg_id, char *ctl, char *cmd, char *trace);
        void ExecuteAsync(int be_id, int msg_id, char *ctl, char *cmd, char *trace);
    private:
        unique_ptr<RemoteExec::Stub> stub_;
        ThreadPool<100> threadpool;
};

class RpcWorker {
    private:
        NetLayer sciNet;
        RemoteExecServiceImpl service;
        FrontBack *rpcClient;

    public:
        RpcWorker();
        ~RpcWorker();

        void runServer();
        FrontBack *getClient() { return rpcClient; }
};

#endif