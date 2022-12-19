# cloudIPtoDB


# Querying Database

1. Get the number of records held in total

```
select count(*) from net;
```

2. Get the number of records from AWS

```
select count(*) from net where cloudplatform='aws';
```
3. Find if a specific IP address exists
For each of record of CIDR records in the datbase we keep the start and end ip in decimal format.  We can then query if a particular IP address is contained.

Howver, you do need to convert your IP address to an integer before running a query for example:

If you want to know if the IP address 
```
177.71.207.129
```

You would convert that to a decimal 2974273409 then perform the following query:

```
SELECT cloudplatform, net 
FROM net 
WHERE start_ip <= '2974273409'
AND end_ip >= '2974273409';
```