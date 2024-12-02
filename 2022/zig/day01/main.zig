const std = @import("std");
const fs = std.fs;
const print = std.debug.print;

const AocRootDir = "../../..";
const Year = "2022";
const Day = "day01";

fn readPuzzleInput(allocator: std.mem.Allocator) ![]const u8 {
    var path_parts = [_][]const u8{ AocRootDir, Year, "inputs", Day, "input.txt" };
    const filepath = try fs.path.join(allocator, &path_parts);
    defer allocator.free(filepath);

    const file = try fs.cwd().openFile(filepath, .{});
    defer file.close();

    const stat = try file.stat();
    const buf = try file.readToEndAlloc(allocator, stat.size);

    return buf;
}

pub fn part1(lines: *std.mem.SplitIterator(u8, .sequence)) !u32 {
    const allocator = std.heap.page_allocator;
    var sums = std.ArrayList(u32).init(allocator);
    defer sums.deinit();

    var sum: u32 = 0;
    while (lines.next()) |line| {
        // ...
        if (line.len == 0) {
            try sums.append(sum); // done with this elf
            sum = 0; // start new elf count
            continue;
        }
        const calories = try std.fmt.parseInt(u32, line, 10);
        sum = sum + calories;
    }
    try sums.append(sum); // last elf
    return std.mem.max(u32, sums.items);
}

pub fn part2(lines: *std.mem.SplitIterator(u8, .sequence)) !u32 {
    const allocator = std.heap.page_allocator;
    var sums = std.ArrayList(u32).init(allocator);
    defer sums.deinit();

    var sum: u32 = 0;
    while (lines.next()) |line| {
        // ...
        if (line.len == 0) {
            try sums.append(sum); // done with this elf
            sum = 0; // start new elf count
            continue;
        }
        const calories = try std.fmt.parseInt(u32, line, 10);
        sum = sum + calories;
    }
    try sums.append(sum); // last elf
    std.mem.sort(u32, sums.items, {}, std.sort.desc(u32));
    return sums.items[0] + sums.items[1] + sums.items[2];
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const buf = try readPuzzleInput(allocator);
    defer allocator.free(buf);
    var lines1 = std.mem.split(u8, buf, "\n");

    const part1_solution = try part1(&lines1);
    print("--- PART 1 ---\n", .{});
    print("{d}\n", .{part1_solution});

    var lines2 = std.mem.split(u8, buf, "\n");
    const part2_solution = try part2(&lines2);
    print("--- PART 2 ---\n", .{});
    print("{d}\n", .{part2_solution});
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
