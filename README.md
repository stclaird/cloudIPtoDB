# Cloud IP to DB
This project creates a SQLite database containing CIDR records from large Cloud Providers platforms. Currently, the project imports IP CIDR networks from the following:

- Amazon Web Services (AWS)
- Google Cloud (GCP)
- Digital Ocean (DO)
- Microsoft Azure (Azure)
- CloudFlare
- Oracle Cloud

# Technology Stack

The two core elements of this project are:
 - A binary written in GoLang which creates a SQLite database object and populates it with IP CIDR data from various Cloud platform providers.
 - A SQLite database file output containing the Cloud platform providers CIDR information.

The SQLite database schema made up of a single 'net' table

```CREATE TABLE IF NOT EXISTS net (
 	net_id INTEGER PRIMARY KEY,
 	net TEXT NOT NULL,
 	start_ip INT NOT NULL,
 	end_ip INT NOT NULL,
 	url TEXT NOT NULL,
 	cloudplatform TEXT NOT NULL,
 	iptype TEXT NOT NULL
 	);
 ```

# Querying the Database
The following section gives examples on how to query the database using SQL.

1. To get the total number of CIDR records held in the database:

```
select count(*) from net;
```

2. To get the number of CIDR records held in the database belong to the cloud platform AWS.

```
select count(*) from net where cloudplatform='aws';
```
3. Find if a specific IP address exists in one of the cidrs held in the database.

The records in the database are in CIDR network format, and not unpacked into individual IP addresses. Not having individual IP addresses stored as records will make querying for IPs using SQL difficult.

To remedy this and allow for the querying of individual IP addresses we also store along side the CIDR record, the start (network) and end (broadcast) address. They are both stored as integers and this means we are able to query whether a specific IP address record exists in the database by testing if the IP address falls between the start record and the end record.

One thing though, for this to work you do need to convert your IP address to an integer before running a query. 
For example, if you want to know if the IP address `177.71.207.129` is withi one of the CIDR records stored in the database:

Firstly you need to convert this IPv4 address to a decimal integer, which is 2974273409 and then perform the following query:

```
SELECT cloudplatform, net 
FROM net 
WHERE start_ip <= '2974273409'
AND end_ip >= '2974273409';
```
If this IP address is contained within one of the CIDR records this return the CIDR record, otherwise, if the IP address is not stored then the database will return no records.

