#include <stdio.h>
#include <stdlib.h>

typedef enum { false, true } bool;

bool has_left(int totalRow, int totalCol, int curRow, int curCol) {
    if (curCol > 0) {
        return true;
    }
    return false;
}

bool has_up(int totalRow, int totalCol, int curRow, int curCol) {
    if (curRow > 0) {
        return true;
    }
    return false;
}

bool has_right(int totalRow, int totalCol, int curRow, int curCol) {
    if (curCol < totalCol-1) {
        return true;
    }
    return false;
}

bool has_down(int totalRow, int totalCol, int curRow, int curCol) {
    if (curRow < totalRow-1) {
        return true;
    }
    return false;
}

bool search_from_pos(char **borad,
                    int boardRowSize,
                    int boardColSize,
                    int curRow,
                    int curCol,
                    char *word,
                    bool **matchMarkPtr) {
    if (*word == '\0') {
        return true;
    }

    if (borad[curRow][curCol] == *word && matchMarkPtr[curRow][curCol] != true) {
        // 这里要特殊处理边界，不优雅
        if (boardRowSize == 1 && boardColSize == 1 && *(word+1) == '\0') {
            return true;
        }
        matchMarkPtr[curRow][curCol] = true;
        if (has_up(boardRowSize, boardColSize, curRow, curCol) == true && \
                search_from_pos(borad, boardRowSize, boardColSize, curRow-1, curCol, word+1, matchMarkPtr) == true) {
            return true;
        }
        if (has_right(boardRowSize, boardColSize, curRow, curCol) == true && \
                search_from_pos(borad, boardRowSize, boardColSize, curRow, curCol+1, word+1, matchMarkPtr) == true) {
            return true;
        }
        if (has_down(boardRowSize, boardColSize, curRow, curCol) == true && \
                search_from_pos(borad, boardRowSize, boardColSize, curRow+1, curCol, word+1, matchMarkPtr) == true) {
            return true;
        }
        if (has_left(boardRowSize, boardColSize, curRow, curCol) == true && \
                search_from_pos(borad, boardRowSize, boardColSize, curRow, curCol-1, word+1, matchMarkPtr) == true) {
            return true;
        }
        matchMarkPtr[curRow][curCol] = false;
    } 

    return false;
}

bool exist(char** board, int boardRowSize, int boardColSize, char* word) {
    int curRow = 0;
    int curCol = 0;

    bool *matchMark = (bool *)calloc(boardRowSize*boardColSize, sizeof(bool));
    bool **matchMarkPtr = (bool **)calloc(boardRowSize, sizeof(bool *));
    for (int i = 0; i < boardRowSize; i++) {
        // 注意这里的两种类型的转换
        matchMarkPtr[i] = &matchMark[i*boardColSize];
    }

    for (int i = 0; i < boardRowSize; ++i) {
        for (int j = 0; j < boardColSize; ++j) {
            if (search_from_pos(board, boardRowSize, boardColSize, curRow, curCol, word, matchMarkPtr) == true) {
                return true;
            }
        }
    }

    free(matchMark);
    free(matchMarkPtr);
    return false;
}

int main() {
    /* char b1[] = {'A','B','C','E'}; */
    /* char b2[] = {'S','F','C','S'}; */
    /* char b3[] = {'A','D','E','E'}; */
    /* char *board[3] = {b1,b2,b3}; */
    // 指向指针的指针和多维数据是类型不兼容的
    char b1[] = {'A'};
    char *board[1] = {b1};
    char *key;
    bool result;

    /* key = "ABCCED"; */
    /* result = exist(board, 3, 4, key); */
    /* printf("%s %d\n", key, result); */

    /* key = "SEE"; */
    /* result = exist(board, 3, 4, key); */
    /* printf("%s %d\n", key, result); */

    /* key = "ABCB"; */
    /* result = exist(board, 3, 4, key); */
    /* printf("%s %d\n", key, result); */

    key = "A";
    result = exist(board, 1, 1, key);
    printf("%s %d\n", key, result);
    key = "AB";
    result = exist(board, 1, 1, key);
    printf("%s %d\n", key, result);
}
