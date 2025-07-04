-- An implementation of the Modifier pattern where immutable values can have functions applied through a chain
-- with original value preserved and changes computed only when accessed
-- written for lua 5.4

Value = {}
Value.__index = Value

function Value:new(val)
    local obj = {
        value = val,
        modifiers = {}
    }
    setmetatable(obj, self)
    return obj
end

function Value:get()
    local result = self.value
    for _, modifier in ipairs(self.modifiers) do
        result = modifier.func(result)
    end
    return result
end

function Value:get_original()
    return self.value
end

function Value:all_modifiers()
    local names = {}
    for _, modifier in ipairs(self.modifiers) do
        table.insert(names, modifier.name)
    end
    return names
end

Modifier = {}
Modifier.__index = Modifier

function Modifier:new(name, func)
    local obj = {
        name = name,
        func = func
    }
    setmetatable(obj, self)
    return obj
end

function Modifier:apply(value)
    table.insert(value.modifiers, self)
    return value
end

local myValue = Value:new(10)
print("Initial value:", myValue:get())

local doubler = Modifier:new("doubler", function(x) return x * 2 end)
local adder = Modifier:new("adder", function(x) return x + 5 end)
local squarer = Modifier:new("squarer", function(x) return x * x end)

doubler:apply(myValue)
adder:apply(myValue)
squarer:apply(myValue)

local anotherValue = Value:new(3)
local multiplier = Modifier:new("multiplier", function(x) return x * 10 end)
local subtractor = Modifier:new("subtractor", function(x) return x - 7 end)

multiplier:apply(anotherValue)
subtractor:apply(anotherValue)

local chainedValue = Value:new(2)
Modifier:new("add_one", function(x) return x + 1 end):apply(chainedValue)
Modifier:new("triple", function(x) return x * 3 end):apply(chainedValue)
Modifier:new("subtract_two", function(x) return x - 2 end):apply(chainedValue)
print("Chained value result:", chainedValue:get())

print("Applied modifiers to myValue:", table.concat(myValue:all_modifiers(), ", "))
print("Applied modifiers to anotherValue:", table.concat(anotherValue:all_modifiers(), ", "))
print("Applied modifiers to chainedValue:", table.concat(chainedValue:all_modifiers(), ", "))
print("Original value of myValue:", myValue:get_original())
print("Modified value of myValue:", myValue:get())

print("myValue original:", myValue:get_original(), "modified:", myValue:get())
print("anotherValue original:", anotherValue:get_original(), "modified:", anotherValue:get())
print("chainedValue original:", chainedValue:get_original(), "modified:", chainedValue:get())
