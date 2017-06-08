#include <stdio.h>
#include <Windows.h>
#include <stdlib.h>
#include <string.h>

#define BUFFER_SIZE 512

//MAC TIME 분석 프로그램(?) MAC TIME 변경부분은 액세스 거부 문제로 아직 미완성


//dump code -[해킹]공격의 예술- dumpcode 참조/
void dump(const unsigned char *data_buffer, const unsigned int length) {
	unsigned char byte;
	unsigned int i, j;
	for(i=0; i < length; i++) {
		byte = data_buffer[i];
		printf("%02x ", data_buffer[i]);  // display byte in hex
		if(((i%16)==15) || (i==length-1)) {
			for(j=0; j < 15-(i%16); j++)
				printf("   ");
			printf("| ");
			for(j=(i-(i%16)); j <= i; j++) {  // display printable bytes from line
				byte = data_buffer[j];
				if((byte > 31) && (byte < 127)) // outside printable char range
					printf("%c", byte);
				else
					printf(".");
			}
			printf("\n"); // end of the dump line (each line 16 bytes)
		} // end if
	} // end for
}

// Read from sector //
int main(void)
{
    FILE *volume;
    int k = 0;
    int count=0;
    long long sector = 0;
    unsigned char buf[BUFFER_SIZE*4] = {0};
	unsigned char reserve_count[4]={0};
	unsigned char nof[4]={0};
	unsigned char fs[4]={0};
	int _reserve=0;
	int _nof=0;
	int _fs=0;
	int move=0;
	int total=0;
 
    sector = 0;             // Sector-aligned offset
 
    volume = fopen("\\\\.\\D:", "r");
    setbuf(volume, NULL);       // Disable buffering
    if(!volume)
    {
        printf("드라이브를 열수 없습니다.\n");
        return 1;
    }
 
    if(fseek(volume, sector*BUFFER_SIZE, SEEK_SET) != 0)
    {
        printf("해당 섹터로 이동불가\n");
        return 2;
    }
 
    // read what is in sector and put in buf //
    fread(buf, sizeof(*buf), BUFFER_SIZE, volume);
	printf("-----------부트레코드 %d sector -----------\n",sector);

	
 
    // Print out what wiat in sector //
	//dump(buf,BUFFER_SIZE);
    for(k=0;k<BUFFER_SIZE;k++){
		if(k==14){
			reserve_count[0]=buf[k];
		}
		if(k==15){
			reserve_count[1]=buf[k];
		}
		if(k==16){
			nof[0]=buf[k];
		}
		if(k==36){
			fs[0]=buf[k];
		}
		if(k==37){
			fs[1]=buf[k];
		}
		if(k==38){
			fs[2]=buf[k];
		}
		if(k==39){
			fs[3]=buf[k];
		}
        printf("%02x ", *(&buf[0]+k));
			count++;
			if(count==16){
				count=0;
				printf("\n");
			}
		}

	printf("%02x %02x , %02x , %02x %02x %02x %02x\n",reserve_count[0],reserve_count[1],nof[0],fs[0],fs[1],fs[2],fs[3]);
	_reserve=*(int *)reserve_count;
	_nof=*(int *)nof;
	_fs=*(int *)fs;
	printf("%d %d %d\n",_reserve,_nof,_fs);
	move=_nof*_fs;
	total=(_reserve*512)+(move*512);
	
	printf("BR(Boot record) sector + %d = Root directory entry sector\n",total/512);

	if(fseek(volume, total, SEEK_SET) != 0)
    {
        printf("해당 섹터로 이동불가\n");
        return 2;
    }

	 fread(buf, sizeof(*buf), BUFFER_SIZE*5, volume);
	 printf("-------------루트 디렉토리 엔트리(MAC TIME정보가 있음) sector %d-----------\n",total/512);
	 dump(buf,BUFFER_SIZE*5);
	fclose(volume);
	system("pause");
 
    return 0;
}