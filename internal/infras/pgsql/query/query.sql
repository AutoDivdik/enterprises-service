-- name: GetAll :many

SELECT e.id,
       e.name,
       country,
       maintenanceYear,
       phone,
       fax,
       t.id   as "type_of_ownership_id",
       t.name as "type_of_ownership_name"
FROM "enterprises".enterprises e
         JOIN "enterprises".types_of_ownership t ON e.type_of_ownership_id = t.id;

-- name: GetEnterpriseByID :one

SELECT e.id,
       e.name,
       country,
       maintenanceYear,
       phone,
       fax,
       t.id   as "type_of_ownership_id",
       t.name as "type_of_ownership_name"
FROM "enterprises".enterprises e
         JOIN "enterprises".types_of_ownership t ON e.type_of_ownership_id = t.id
WHERE e.id = $1 LIMIT 1;