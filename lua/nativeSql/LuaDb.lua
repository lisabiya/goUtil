---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by wakfu.
--- DateTime: 2020/8/1 11:37
---
LuaDb = {}

sqlite3 = require('sqlite3')

function LuaDb.new(tableName)
    local obj = sqlite3.new(tableName)
    return obj
end


--*********************go实现的接口api**************************

---@return boolean|table
function LuaDb:open(dbName, table)

end

function LuaDb:query(sqlStr)
end


return LuaDb