local tools = require("lua/tools")
local json = require("lua.json")

function testGoFunc()
    local params = { GO_Util.printMe("来自lua的问候", "lehao ") }
    for k, v in pairs(params) do
        print(k, v)
    end
end

function testIt(a, b)
    return a + b, 100
end

function initParams()
    local param = getParams("name")

    sqlite3 = require('sqlite3')
    c = sqlite3.new()
    ok, err = c:open('salary.db', { cache = 'shared' })
    if ok then

        local key, value = tools.sqlInsert({ name = "小明", department = "实习",
                                               social_security = 100 })
        res, err = c:query("INSERT INTO t_salary(department,name,social_security) values('实习','小明5','1002')")
        --res, err = c:query("select * from t_salary")
        return { res, err, "sd" }
    end

    return "失败"
end

function testSql()
    local key, value = tools.sqlInsert({ name = "小明", department = "实习",
                                         social_security = 100 })
    print("INSERT INTO t_salary(" .. key .. ") values(" .. value .. ")")
end

testSql()