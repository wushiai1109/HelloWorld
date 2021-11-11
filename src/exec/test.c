//把文件读到内存
#include<stdio.h>
#include<stdlib.h>
#include<string.h>
#define maxCount 1024

void readtoMem(char ***pwd,char *path){
    FILE *fp = fopen(path,"r+");
    //得到文件有几行 
    int lineCount;
    char buf[maxCount]; 
    while(fgets(buf,maxCount,fp)){
        lineCount++;
    }
    rewind(fp);//重新返回文件开头 
    *pwd = (char **)malloc((lineCount+1)*sizeof(char *));
    char **t = *pwd;
    while(fgets(buf,maxCount,fp)){
        int n = strlen(buf);
        *t = (char*)malloc((n+1)*sizeof(char *));
        strcpy(*t,buf);
        t++;
    }
    t = NULL;
}
int main(void){
    char** pwd;
    readtoMem(&pwd,"test1.txt") ;
    while(*pwd){
        printf("%s",*pwd++);
    }
    return 0;
}