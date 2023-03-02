

Create entries in Employee collection under local db: 

db.getCollection("Employee").insertMany([
 { id: "1", name: "Emp-1", age: 25, designation: "developer" },
 { id: "2", name: "Emp-2", age: 26, designation: "sr. developer" },
 { id: "3", name: "Emp-3", age: 27, designation: "manager" }
 ])
