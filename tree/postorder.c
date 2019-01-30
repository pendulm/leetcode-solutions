#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};


int *postorderTraversal(struct TreeNode *root, int *returnSize) {
    if (!root) {
        *returnSize = 0;
        return NULL;
    }

    int lsize = 0, rsize = 0;
    int *larr = NULL, *rarr = NULL;

    if (root->left) {
        larr = postorderTraversal(root->left, &lsize);
    }

    if (root->right) {
        rarr = postorderTraversal(root->right, &rsize);
    }

    *returnSize = 1 + lsize + rsize;
    int *arr = calloc(*returnSize, sizeof(int));

    if (!arr) {
        exit(1);
    }

    int i = 0;
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
    arr[i++] = root->val;

    return arr;
}




int main() {
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

    int size = 0;
    int *arr = postorderTraversal(&l1, &size);

    for (int i = 0; i < size; i++) {
        printf("%d\n", arr[i]);
    }

    if (arr) {
        free(arr);
    }

    return 0;
}
