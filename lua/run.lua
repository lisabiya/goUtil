function testGoFunc()
    local params = { GO_Util.printMe("来自lua的问候", "lehao ") }
    for k, v in pairs(params) do
        print(k, v)
    end
end

function initParams()
    local example = require('lua.db_module.example')
    local code, tables = example.update()

    return { code = code, response = tables }
end

local Builder = require "lua.LuaQuB"

local object = Builder.new()
                      :update("t_salary", { name = "李雷", department = "油烟清理", social_security = 1500 })
                      :where("`id` =", "10")
print(tostring(object))