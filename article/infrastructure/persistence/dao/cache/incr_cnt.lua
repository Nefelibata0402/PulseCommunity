-- 具体业务 并加上id
local key = KEYS[1]
-- 是阅读数，点赞数还是收藏数 并加上id
local cntKey = ARGV[1]
--加一还是减一
local delta = tonumber(ARGV[2])

local exist=redis.call("EXISTS", key)
if exist == 1 then
    redis.call("HINCRBY", key, cntKey, delta)
    return 1
else
    return 0
end