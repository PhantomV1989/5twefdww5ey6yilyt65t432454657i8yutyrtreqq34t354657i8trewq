# Design considerations
1. HTTP is used here because the network layer is expected to handle SSL (eg. load balancer, proxies like nginx)
2. Request/Response format may be slightly different from assignment to keep consistent because there are some inconsistent formats used in assignment (eg. some formats are nested jsons in the data field) 
3. Hardcoded image limit of 100 since nothing(pagnation? size?) is specified in the assignment
4. DB should implement some form of access control but for convenience sake, root is used
