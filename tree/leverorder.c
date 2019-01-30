#include <stdio.h>
#include <stdlib.h>

#define QLEN 100

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

struct Queue {
    struct TreeNode **q;
    int head, tail;
    int len;
};

struct Queue *createQueue() {
    struct Queue *ptr;

    ptr = calloc(1, sizeof(struct Queue));
    ptr->q = calloc(QLEN, sizeof(struct TreeNode *));
    ptr->len = 0;
    ptr->head = ptr->tail = 0;
    return  ptr;
}

int destroyQueue(struct Queue *ptr) {
    if (ptr) {
        free(ptr->q);
        free(ptr);
        ptr = NULL;
        return 1;
    }
    return 0;
}

int enQueue(struct Queue *queue, struct TreeNode *t) {
    int st = 0;
    if ((queue->tail - queue->head + QLEN) % QLEN == QLEN-1) {
        queue->q[queue->tail] = t;
        queue->tail = (queue->tail + 1) % QLEN;
    } else {
        // full
        st = 1;
    }
    return st;
}

struct TreeNode *deQueue(struct Queue *queue) {
    struct TreeNode *t = NULL;
    if (queue->tail != queue -> head) {
        t = queue->q[queue->head];
        queue->head = (queue->head + 1) % QLEN;
    }

    return t;
}

int doBFS(struct Queue *qptr) {
    int st = 1;
    struct TreeNode *tp = deQueue(qptr);
    if (tp) {
        if (tp->left) {
            enQueue(qptr, tp->left);
        }
        if (tp->right) {
            enQueue(qptr, tp->right);
        }
        doBFS(qptr);
    } else {
        st = 0;
    }
    return st;
}




int** levelOrder(struct TreeNode* root, int** columnSizes, int* returnSize) {
    if (!root) {
        *columnSizes = NULL;
        *returnSize = 0;
        return NULL;
    }

    struct Queue *qptr = createQueue();

    enQueue(qptr, root);

    doBFS(queue);
}
    
