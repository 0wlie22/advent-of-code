#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int compare(const void *a, const void *b) { return (*(int *)a - *(int *)b); }

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
  int *list_one = (int *)malloc(array_size * sizeof(int));
  int *list_two = (int *)malloc(array_size * sizeof(int));
  int num_elements = 0;
  int value;

  while ((read = getline(&line, &len, fptr)) != -1) {
    first_token = atoi(strtok(line, delim));
    if (first_token != 0) {
      second_token = atoi(strtok(NULL, delim));
    }

    if (num_elements >= array_size) {
      array_size *= 2;
      list_one = (int *)realloc(list_one, array_size * sizeof(int));
      list_two = (int *)realloc(list_two, array_size * sizeof(int));
    }

    list_one[num_elements] = first_token;
    list_two[num_elements] = second_token;
    num_elements++;
  }

  qsort(list_one, num_elements, sizeof(int), compare);
  qsort(list_two, num_elements, sizeof(int), compare);

  int sum = 0;
  for (int i = 0; i < num_elements; i++) {
    sum += abs(list_one[i] - list_two[i]);
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
