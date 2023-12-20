window_title = "Audio Example"

function at_fire_pressed()
  print("playing sound")
  playsound(440, 100)
  print("done playing sound")
end

function at_esc_pressed()
  quit()
end
