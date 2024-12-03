#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int compare(const void *a, const void *b) { return (*(int *)a - *(int *)b); }

typedef struct data {
  int token;
  int count;
} tData;

tData *linear_search(tData *arr, int size, int key) {
  for (int i = 0; i < size; i++) {
    if (arr[i].token == key) {
      return &arr[i];
    }
  }
  return NULL;
}

int main(int argc, char *argv[]) {
  char *filename = argv[1];
  FILE *fptr;

  char *line = NULL;
  size_t len = 0;
  ssize_t read;

  char *delim = "   ";
  int first_token;
  int second_token;

  fptr = fopen(filename, "r");
  if (fptr == NULL) {
    printf("Cannot open file \n");
    exit(0);
  }

  int array_size = 2;
  tData *list_one = (tData *)malloc(array_size * sizeof(tData));
  tData *list_two = (tData *)malloc(array_size * sizeof(tData));
  int num_elements = 0;

  while ((read = getline(&line, &len, fptr)) != -1) {
    first_token = atoi(strtok(line, delim));
    if (first_token != 0) {
      second_token = atoi(strtok(NULL, delim));
    }

    if (num_elements >= array_size) {
      array_size *= 2;
      list_one = (tData *)realloc(list_one, array_size * sizeof(tData));
      list_two = (tData *)realloc(list_two, array_size * sizeof(tData));
    }

    if (linear_search(list_one, num_elements, first_token) == NULL) {
      list_one[num_elements].token = first_token;
      list_one[num_elements].count = 1;
    } else {
      list_one[linear_search(list_one, num_elements, first_token) - list_one]
          .count++;
    }

    if (linear_search(list_two, num_elements, second_token) == NULL) {
      list_two[num_elements].token = second_token;
      list_two[num_elements].count = 1;
    } else {
      list_two[linear_search(list_two, num_elements, second_token) - list_two]
          .count++;
    }

    num_elements++;
  }

  int sum = 0;
  for (int i = 0; i < num_elements; i++) {
    tData *found = linear_search(list_two, num_elements, list_one[i].token);
    if (found != NULL) {
      sum += list_one[i].token * list_one[i].count * found->count;
    }
  }
  printf("%d\n", sum);

  fclose(fptr);
  if (line) {
    free(line);
  }

  free(list_one);
  free(list_two);

  return 0;
}
