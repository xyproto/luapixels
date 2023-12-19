window_title = "Pixel McPixelface"

cursor_x = 100
cursor_y = 100

function at_start()
    print("hi")
end

function at_every_frame()
    background(0, 128, 255)
    for y = 0, 199 do
        for x = 0, 255 do
            plot(x, y, x)
        end
        for x = 256, 318 do
            plot(x, y, 50)
        end
        plot(319, y, 70)
    end
    plot(318, 198, 41)
    plot(318, 199, 41)
    plot(319, 198, 41)
    plot(319, 199, 41)
    plot(cursor_x, cursor_y, 41)
end

function at_keypress()
    local key = last_key
    -- Quit if GLFW key codes 81 ('q') or 256 ('Esc') are pressed
    if key == 81 or key == 256 then
        quit()
    end

    -- Adjust cursor position based on arrow keys or WASD
    -- GLFW key codes: 87 ('W'), 65 ('A'), 83 ('S'), 68 ('D')
    -- Arrow keys: 262 (Right), 263 (Left), 264 (Down), 265 (Up)
    if key == 87 or key == 265 then
        cursor_y = cursor_y - 1
    elseif key == 83 or key == 264 then
        cursor_y = cursor_y + 1
    elseif key == 65 or key == 263 then
        cursor_x = cursor_x - 1
    elseif key == 68 or key == 262 then
        cursor_x = cursor_x + 1
    end
end

function at_end()
    print("bye!")
end
