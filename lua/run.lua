local luaDbSqLite = require('lua.db_module.db_module')

function testGoFunc()
    local params = { GO_Util.printMe("来自lua的问候", "lehao ") }
    for k, v in pairs(params) do
        print(k, v)
    end
end


function initParams()
    local ormDb = luaDbSqLite.new("t_salary")
    print(ormDb:Tag())
    local code, tables = ormDb:Rows()
    print(ormDb:Tag())
    ormDb:Insert({ name = "小明", department = "实习",
                   social_security = 100 })
    print(ormDb:Tag())
    return { code = code, response = tables }
end

