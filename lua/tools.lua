---
--- Generated by EmmyLua(https://github.com/EmmyLua)
--- Created by wakfu.
--- DateTime: 2020/6/9 10:58
---
local tools = { _version = "1.0" }

--*******************************工具栏*******************************
---文件是否存在
---@param path string 文件名
function tools.file_exists(path)
    local file = io.open(path, "r")
    if file then
        file:close()
    end
    return file ~= nil
end

-- trim
---@param str string
----@return string
function tools.trimAll(str)
    if type(str) == "number" then
        return false
    end
    return str:gsub("^%s*(.-)%s*$", "%1")
end

-- 查询是否为空字符
---@param str string
----@return boolean
function tools.isStrEmpty(str)
    if str then
        --前后trim空格
        local strNew = trimAll(str)
        if strNew == "" or strNew == "''" then
            return true
        end
        return false
    else
        return true
    end
end

--是否是ip地址
---@param ipStr string
---@return boolean 是否是ip地址
function tools.isIP(ipStr)
    if type(ipStr) == 'string' then
        local ips = { ipStr:find("^(%d+)%.(%d+)%.(%d+)%.(%d+)$") }
        if #ips > 0 then
            return true
        end
        return false
    else
        return false
    end
end


-- sql 语句生成
---@param table table  [string,string]
---@return string,string
function tools.sqlInsert(table)
    local key, value = "", ""
    for k, v in pairs(table) do
        key = key .. k .. ","
        value = value .. v .. ","
    end
    return key, value
end

return tools