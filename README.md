# cloudIPtoDB


# Querying Database

1. Get the number of CIDR records held in total

```
select count(*) from net;
```

2. Get the number of CIDR records held in the database belong to the cloud platform AWS.

```
select count(*) from net where cloudplatform='aws';
```
3. Find if a specific IP address exists in one of the cidrs held in the database.

The records in the database are in CIDR network format, and not unpacked into individual IP addresses. Not having individual IP addresses stored as records will make querying for IPs using SQL difficult.

Toremedy this and allow for the querying of individual IP addresses we also store along side the CIDR record, the start (network) and end (broadcast) address. The are both stored as integers and this means we are able to query wether a specific IP address record exists in the database by testing if the IP address falls between the start record and the end record.

One thing though, for this to work you do need to convert your IP address to an integer before running a query. 
For example, if you want to know if the IP address:
```
177.71.207.129
```

Firstly you need to convert this IPv4 address to a decimal integer, which is 2974273409 and then perform the following query:

```
SELECT cloudplatform, net 
FROM net 
WHERE start_ip <= '2974273409'
AND end_ip >= '2974273409';

If this IP address is contained within one of the CIDR records this return the CIDR record, other wise, if the IP address is not stored then the database will return no records.
```
