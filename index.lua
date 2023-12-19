windowTitle = "Pixel McPixelface"

function atStart()
    -- Set palette color number 41 to red (255, 0, 0)
    setpal(41, 255, 0, 0)
end

function everyFrame()
    background(0, 128, 255)
    plot(0, 0, 41)
    line(0, 0, 319, 199, 41)
end

function atEnd()
    -- Print a friendly message to the console
    print("bye!")
end
