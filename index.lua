windowTitle = "Pixel McPixelface"

function atStart()
    -- Set palette color number 41 to red (255, 0, 0)
    setpal(41, 255, 0, 0)
end

function atEveryFrame()
    background(0, 128, 255)
    -- Fill the screen with color indexes from left to right
    for y = 0, 199 do
        for x = 0, 255 do
            plot(x, y, x)
        end
        for x = 256, 318 do
            plot(x, y, 28)
        end
        plot(319, y, 41)
    end
end

function atEnd()
    -- Print a friendly message to the console
    print("bye!")
end
