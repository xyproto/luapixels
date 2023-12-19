window_title = "Double Jump"

player_x = 100
player_y = 100

player_ax = 0
player_ay = 4

earth_ax = 0
earth_ay = 1

up_pressed = false
down_pressed = false
left_pressed = false
right_pressed = false
space_pressed = false

function at_start()
    print("Double Jump")
end

function at_every_tick()

    -- Apply friction if player_y == 199 and neither left nor right arrow key is held down
    if player_y == 199 and not left_pressed and not right_pressed then
        player_ax = player_ax * 0.7
    elseif player_y == 199 and left_pressed or right_pressed then
        player_ax = player_ax * 0.7
    end
    if up_pressed then
        player_ay = player_ay - 5
        player_y = player_y - 1
    end
    if down_pressed then
        player_ay = player_ay + 5
        player_y = player_y + 1
    end
    if left_pressed then
        player_ax = player_ax - 5
        player_x = player_x - 1
    end
    if right_pressed then
        player_ax = player_ax + 5
        player_x = player_x + 1
    end
    if space_pressed then
        player_ay = -10
        -- Do not support sustained jumping
        space_pressed = false
    end

    -- Gravity
    player_ax = player_ax + earth_ax
    player_ay = player_ay + earth_ay

    -- Player movement
    player_x = player_x + player_ax
    player_y = player_y + player_ay

    -- Logic for holding down keys
    if player_y < 199 then
        up_pressed = false
        down_pressed = false
        left_pressed = false
        right_pressed = false
    end

    if player_y > 199 then
        player_y = 199
        player_ay = -player_ay
        player_ay = player_ay * 0.5
    elseif player_y < 0 then
        player_y = 0
        player_ay = -player_ay
        player_ay = player_ay * 0.5
    end
    if player_x > 319 then
        player_x = 319
        player_ax = -player_ax
        player_ax = player_ax * 0.5
    elseif player_x < 0 then
        player_x = 0
        player_ax = -player_ax
        player_ax = player_ax * 0.5
    end
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
    plot(player_x, player_y, 41)
end

function at_direction_pressed(x, y)
    if x == -1 then left_pressed = true end
    if x == 1 then right_pressed = true end
    if y == -1 then down_pressed = true end
    if y == 1 then up_pressed = true end
end

function at_direction_released(x, y)
    if x == -1 then left_pressed = false end
    if x == 1 then right_pressed = false end
    if y == -1 then down_pressed = false end
    if y == 1 then up_pressed = false end
end

function at_keypress()
    local key = last_key
    -- Quit if GLFW key codes 81 ('q') or 256 ('Esc') are pressed
    if key == 81 or key == 256 then
        quit()
    end
    -- Jump logic
    if key == 32 then space_pressed = true end
end

function at_keyrelease()
    local key = last_key
    -- Disable jump on key release
    if key == 32 then space_pressed = false end
end

function at_end()
    print("bye!")
end
