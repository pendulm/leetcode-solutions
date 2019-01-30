awk '
{
    n=1;
    while(n <= NF) {
        mem[NR, n] = $n;
        n++
    }
};
END {
col=1;
    while(col<n) {
        row=1;
        while (row <= NR) {
            printf("%s", mem[row,col]);
            if (row != NR) 
                printf(" ");
                row++
        };
        if (col != n)
            printf("\n");
        col++
    } 
}' file.txt
