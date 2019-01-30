#include <malloc.h>
#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
     int val;
     struct TreeNode *left;
     struct TreeNode *right;
};

int* preorderTraversal(struct TreeNode* root, int* returnSize) {
    if (!root) {
        *returnSize = 0;
        return NULL;
    }

    int lsize = 0, rsize = 0;
    int *larr = NULL, *rarr = NULL;
    if (root->left) {
        larr = preorderTraversal(root->left, &lsize);
    }
    if (root->right) {
        rarr = preorderTraversal(root->right, &rsize);
    }

    *returnSize = lsize + rsize + 1;
    int *arr = (int *)calloc(*returnSize, sizeof(int));

    int i = 0;
    arr[i++] = root->val;

    if (larr) {
        for (int j = 0; j < lsize; j++) {
            arr[i++] = larr[j];
        }
        free(larr);
    }

    if (rarr) {
        for (int j = 0; j < rsize; j++) {
            arr[i++] = rarr[j];
        }
        free(rarr);
    }

    return arr;
}

struct TreeNode *buildTree(int const arr[], int n) {
    if (n == 0) {
        return NULL;
    }

    struct TreeNode *root = NULL;
    root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
    root->val = arr[0];

    if (n > 1) {
        struct TreeNode *left = NULL;
        if (arr[1] != 0) {
            left = (struct TreeNode *)malloc(sizeof(struct TreeNode));
            left->val = arr[1];
        }
        root->left = left;
    }

    if (n > 2) {
        struct TreeNode *right = NULL;
        if (arr[2] != 0) {
            right = (struct TreeNode *)malloc(sizeof(struct TreeNode));
            right->val = arr[2];
        }
        root->right = right;
    }
    return root;
}

int main() {
    // [37,-34,-48,null,-100,-100,48,null,null,null,null,-54,null,-71,-22,null,null,null,8]
    struct TreeNode l6 = {8, NULL, NULL};
    struct TreeNode l5_1 = {-71, NULL, NULL};
    struct TreeNode l5_2 = {-22, NULL, &l6};
    struct TreeNode l4 = {-54, &l5_1, &l5_2};
    struct TreeNode l3_1 = {-100, NULL, NULL};
    struct TreeNode l3_2 = {-100, NULL, NULL};
    struct TreeNode l3_3 = {48, &l4, NULL};
    struct TreeNode l2_1 = {-34, NULL, &l3_1};
    struct TreeNode l2_2 = {-48, &l3_2, &l3_3};
    struct TreeNode l1 = {37, &l2_1, &l2_2};

    /* struct TreeNode l5 = {5, NULL, NULL}; */
    /* struct TreeNode l4 = {2, &l5, NULL}; */
    /* struct TreeNode l3 = {4, &l4, NULL}; */
    /* struct TreeNode l2 = {1, &l3, NULL}; */
    /* struct TreeNode l1 = {3, &l2, NULL}; */

    int *arr = NULL, size = 0;
    arr = preorderTraversal(&l1, &size);

    for (int i = 0; i < size; i++) {
        printf("%d\n", arr[i]);
    }

    if (arr) {
        free(arr);
    }
    return 0;
}
