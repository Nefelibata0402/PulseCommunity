-- 1, 2, 3, 4, 5, 6, 7 这是你的元素
-- ZREMRANGEBYSCORE key1 0 6
-- 7 执行完之后

-- 限流对象
local key = KEYS[1]
-- 窗口大小 1s
local window = tonumber(ARGV[1])
-- 阈值 1个
local threshold = tonumber(ARGV[2])
-- 当前时间
local now = tonumber(ARGV[3])
-- 窗口的起始时间
local min = now - window
-- 删除在窗口起始时间之前的所有记录，确保只保留窗口内的记录。
redis.call('ZREMRANGEBYSCORE', key, '-inf', min)
-- 计算当前窗口内的记录数量，即当前窗口内的请求数量。
local cnt = redis.call('ZCOUNT', key, '-inf', '+inf')
-- 如果当前窗口内的请求数量大于等于阈值（threshold），表示需要执行限流
if cnt >= threshold then
    -- 执行限流
    return "true"
else
    -- 将当前时间作为 score 和 member 添加到有序集合中，表示新的请求。
    redis.call('ZADD', key, now, now)
    -- 设置有序集合的过期时间为窗口大小，确保过期后自动清理无效记录。
    redis.call('PEXPIRE', key, window)
    return "false"
end