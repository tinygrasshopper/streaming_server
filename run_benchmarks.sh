ab -c1 -n100 http://localhost:8080/ > benchmark_1_concurrent.txt
ab -c10 -n100 http://localhost:8080/ > benchmark_10_concurrent.txt
ab -c50 -n1000 http://localhost:8080/ > benchmark_50_concurrent.txt
ab -c100 -n1000 http://localhost:8080/ > benchmark_100_concurrent.txt
