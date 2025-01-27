-- name: GetEntryByDomain :one
select * from safety_domain_blocklist
where domain = $1
limit 0, 1;

-- name: ListDomainBlocklist :many
select * from safety_domain_blocklist;
