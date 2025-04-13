#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/time.h>
#include <unistd.h>

long long get_current_time_ms() {
  struct timeval tv;
  gettimeofday(&tv, NULL);
  return (tv.tv_sec * 1000) + (tv.tv_usec / 1000);
}

void doRequest() {
  printf("Initiated the request.... On %ld\n", pthread_self());
  sleep(3);
  printf("Finished the request On %ld\n", pthread_self());
}

void start() {
  long long start = get_current_time_ms();

  while (1) {
    printf("Rendering UI On... %ld\n", pthread_self());
    long long now = get_current_time_ms();

    if (now - start >= 5000 && now - start < 6000) {
      doRequest();
    }

    sleep(1);
  }
}

int main(int argc, char *argv[]) {
  start();
  return EXIT_SUCCESS;
}
