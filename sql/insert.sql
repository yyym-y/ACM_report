INSERT INTO difficulty VALUES(0, "未定级");
INSERT INTO difficulty VALUES(1, "2400");
INSERT INTO difficulty VALUES(2, "1700");
INSERT INTO difficulty VALUES(3, "1300");
INSERT INTO difficulty VALUES(4, "1100");
INSERT INTO difficulty VALUES(5, "800");
INSERT INTO difficulty VALUES(6, "1800");
INSERT INTO difficulty VALUES(7, "1200");
INSERT INTO difficulty VALUES(8, "1000");
INSERT INTO difficulty VALUES(10, "1400");
INSERT INTO difficulty VALUES(11, "1900");
INSERT INTO type_list VALUES(1, "codeforces");

INSERT INTO problem_list VALUES(1, "C2. Magnitude (Hard Version)", "https://codeforces.com/problemset/problem/1984/C2", 1, "2024-7-7", 2);
INSERT INTO problem_list VALUES(2, "C1. Magnitude (Easy Version)", "https://codeforces.com/problemset/problem/1984/C1", 1, "2024-7-7", 3);
INSERT INTO problem_list VALUES(3, "B. Large Addition", "https://codeforces.com/problemset/problem/1984/B", 1, "2024-7-7", 4);
INSERT INTO problem_list VALUES(4, "A. Strange Splitting", "https://codeforces.com/problemset/problem/1984/A", 1, "2024-7-7", 5);
INSERT INTO problem_list VALUES(5, "E. Manhattan Triangle", "https://codeforces.com/problemset/problem/1979/E", 1, "2024-7-6", 1);
INSERT INTO problem_list VALUES(6, "D. Fixing a Binary String", "https://codeforces.com/problemset/problem/1979/D", 1, "2024-7-5", 6);
INSERT INTO problem_list VALUES(7, "C. Earning on Bets", "https://codeforces.com/problemset/problem/1979/C", 1, "2024-7-4", 7);
INSERT INTO problem_list VALUES(8, "B. XOR Sequences", "https://codeforces.com/problemset/problem/1979/B", 1, "2024-7-4", 8);
INSERT INTO problem_list VALUES(9, "A. Guess the Maximum", "https://codeforces.com/problemset/problem/1979/A", 1, "2024-7-4", 5);
INSERT INTO problem_list VALUES(10, "A. Array Divisibility", "https://codeforces.com/problemset/problem/1983/A", 1, "2024-7-8", 9);
INSERT INTO problem_list VALUES(11, "A. Catch the Coin", "https://codeforces.com/problemset/problem/1989/A", 1, "2024-7-12", 5);
INSERT INTO problem_list VALUES(12, "B. Substring and Subsequence", "https://codeforces.com/problemset/problem/1989/B", 1, "2024-7-12", 7);
INSERT INTO problem_list VALUES(13, "C. Two Movies", "https://codeforces.com/problemset/problem/1989/C", 1, "2024-7-12", 10);
INSERT INTO problem_list VALUES(14, "D. Smithing Skill", "https://codeforces.com/problemset/problem/1989/D", 1, "2024-7-12", 11);

SELECT * from problem_list;
SELECT * FROM descripe_list;
SELECT * FROM solution_list;

DELETE FROM problem_list WHERE problem_id=15;


UPDATE descripe_list SET description="Alice and Bob came up with a rather strange game. They have an array of integers $a_1, a_2,\\ldots, a_n$. Alice chooses a certain integer $k$ and tells it to Bob, then the following happens:\n - Bob chooses two integers $i$ and $j$ ($1 \\le i < j \\le n$), and then finds the maximum among the integers $a_i, a_{i + 1},\\ldots, a_j$;\n - If the obtained maximum is **strictly greater** than $k$, Alice wins, otherwise Bob wins. Help Alice find the maximum $k$ at which she is guaranteed to win." WHERE problem_id = 9;

ALTER TABLE problem_list AUTO_INCREMENT = 15;