function atStart()
    -- Set palette color number 41 to red (255, 0, 0)
    setpalette(41, 255, 0, 0)
end

function everyFrame()
    drawBackground(0, 128, 255)
    plot(0, 0, 41)
end

function atEnd()
    -- Print a friendly message to the console
    print("bye!")
end
