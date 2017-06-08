#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <sys/msg.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/ipc.h>
struct thread_snd{

	int *sockfd;
	int *mspid;
	int *ms_datatype1;
	int *ms_datatype2;

};
typedef struct {
	long data_type;
	char data_buf[1024];
}t_data;
void *_wwrite(void *data);
void *_rread(void *data);
int main(){

	int serversock,clientsock1,clientsock2,clen1,clen2,signal;
	int mspid;
	pthread_t thid1;
	pthread_t thid2;
	pid_t pid;
	int ms_datatype1=1;
	int ms_datatype2=2;
	struct thread_snd ths;
	struct thread_snd ths2;
	struct sockaddr_in serveraddr,clientaddr1,clientaddr2;
	memset(&serveraddr,0x00,sizeof(serveraddr));
	serveraddr.sin_family=AF_INET;
	serveraddr.sin_addr.s_addr=htonl(INADDR_ANY);
	serveraddr.sin_port=htons(7892);
	
	serversock=socket(PF_INET,SOCK_STREAM,IPPROTO_TCP);
	bind(serversock,(struct sockaddr *)&serveraddr,sizeof(serveraddr));
	listen(serversock,5);
	clen1=sizeof(clientaddr1);
	clen2=sizeof(clientaddr2);
	mspid=msgget((key_t)1234,IPC_CREAT|0666);
	msgctl(mspid,IPC_RMID,NULL);
	mspid=msgget((key_t)1234,IPC_CREAT|0666);
	if(mspid==-1)
	 printf("msgget();실패");
	while(1){
		clientsock1=accept(serversock,(struct sockaddr *)&clientaddr1,&clen1);
		clientsock2=accept(serversock,(struct sockaddr *)&clientaddr2,&clen2);
		signal=htonl(1);
		write(clientsock1,&signal,sizeof(signal));
		write(clientsock2,&signal,sizeof(signal));
		pid=fork();
		if(pid==0){
			pid=fork();
			if(pid==0){
				ths.sockfd=&clientsock2;
				ths.mspid=&mspid;
				ths.ms_datatype1=&ms_datatype1;
				ths2.sockfd=&clientsock2;
				ths2.mspid=&mspid;
				ths2.ms_datatype1=&ms_datatype2;
				pthread_create(&thid1,NULL,_wwrite,&ths);
				pthread_create(&thid2,NULL,_rread,&ths2);
				pthread_join(thid1,NULL);
				pthread_join(thid2,NULL);
				//ing,..
				

			}
			else{
				ths.sockfd=&clientsock1;
				ths.mspid=&mspid;
			        ths.ms_datatype1=&ms_datatype1;
				ths2.sockfd=&clientsock1;
				ths2.mspid=&mspid;
				ths2.ms_datatype1=&ms_datatype2;
				pthread_create(&thid1,NULL,_wwrite,&ths2);
				pthread_create(&thid2,NULL,_rread,&ths);
				pthread_join(thid1,NULL);
				pthread_join(thid2,NULL);
			}
		}else{
			ms_datatype1+=2;
			ms_datatype2+=2;
			close(clientsock1);
			close(clientsock2);
		}


	}
	


	return 0;
}

void *_wwrite(void *data){
struct thread_snd *ths2;
t_data _data;
ths2=(struct thread_snd *)data;
int *sockfd=ths2->sockfd;
int *mspid=ths2->mspid;
int *ms_datatype1=ths2->ms_datatype1;
_data.data_type=*ms_datatype1;
char buffer[1024];
while(1){
if(recv(*sockfd,buffer,1024,0)<=0){//클라이언트로부터 데이터를 받음
   //send(*sockfd,"close",1024,0);
   strcpy(_data.data_buf,"close");
   printf("%s\n",_data.data_buf);
   msgsnd(*mspid,&_data,sizeof(t_data)-sizeof(long),0);
   break;
}

strcpy(_data.data_buf,buffer);
printf("%s\n",_data.data_buf);
if(-1==msgsnd(*mspid,&_data,sizeof(t_data)-sizeof(long),IPC_NOWAIT)){
printf("msgsnd()실패\n");
break;
}//메세지큐로 보내는 함수
}

}
void *_rread(void *data){
struct thread_snd *ths2;
t_data _data;
ths2=(struct thread_snd *)data;
int *sockfd=ths2->sockfd;
int *mspid=ths2->mspid;
int *ms_datatype1=ths2->ms_datatype1;
_data.data_type=*ms_datatype1;
char buffer[1024];
while(1){
if(msgrcv(*mspid,&_data,sizeof(t_data)-sizeof(long),*ms_datatype1,0)==-1){//메세지큐로 부터 받는 함수
break;
}
printf("fuck\n",_data.data_buf);
strcpy(buffer,_data.data_buf);
/*if(!strcmp(buffer,"close")){
	break;
}*/
send(*sockfd,buffer,1024,0);
}

}
