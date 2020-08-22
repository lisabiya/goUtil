function testGoFunc()
    local params = { GO_Util.printMe("来自lua的问候", "lehao ") }
    for k, v in pairs(params) do
        print(k, v)
    end
end

function initParams()
    local example = require('lua.db_module.example')
    local code, tables = example.getList()

    return { code = code,count=#tables, response = tables }
end
