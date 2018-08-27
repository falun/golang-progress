package spinner

// Generated with the following javascript:
// const convert = input => Object.keys(input).forEach(k => {
//   const def = input[k]
//   const filteredFrames = def.frames.map(f => f.replace(/\\/g, '\\\\'))
//   const frames = filteredFrames.map(f => `"${f}"`)
//   console.log(`add("${k}", []string{${frames.join(', ')}}, ${def.interval})`)
// })
// applied to https://github.com/sindresorhus/cli-spinners/blob/master/spinners.json

var Animations = map[string]Frameset{}

func init() {
	add := func(n string, f []string, r int) {
		Animations[n] = NewFrameset(n, f, r)
	}

	add("dots", []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}, 80)
	add("dots2", []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}, 80)
	add("dots3", []string{"⠋", "⠙", "⠚", "⠞", "⠖", "⠦", "⠴", "⠲", "⠳", "⠓"}, 80)
	add("dots4", []string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"}, 80)
	add("dots5", []string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"}, 80)
	add("dots6", []string{"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"}, 80)
	add("dots7", []string{"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"}, 80)
	add("dots8", []string{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"}, 80)
	add("dots9", []string{"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"}, 80)
	add("dots10", []string{"⢄", "⢂", "⢁", "⡁", "⡈", "⡐", "⡠"}, 80)
	add("dots11", []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}, 100)
	add("dots12", []string{"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨", "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀"}, 80)
	add("line", []string{"-", "\\", "|", "/"}, 130)
	add("line2", []string{"⠂", "-", "–", "—", "–", "-"}, 100)
	add("pipe", []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}, 100)
	add("simpleDots", []string{".  ", ".. ", "...", "   "}, 400)
	add("simpleDotsScrolling", []string{".  ", ".. ", "...", " ..", "  .", "   "}, 200)
	add("star", []string{"✶", "✸", "✹", "✺", "✹", "✷"}, 70)
	add("star2", []string{"+", "x", "*"}, 80)
	add("flip", []string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"}, 70)
	add("hamburger", []string{"☱", "☲", "☴"}, 100)
	add("growVertical", []string{"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"}, 120)
	add("growHorizontal", []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "▊", "▋", "▌", "▍", "▎"}, 120)
	add("balloon", []string{" ", ".", "o", "O", "@", "*", " "}, 140)
	add("balloon2", []string{".", "o", "O", "°", "O", "o", "."}, 120)
	add("noise", []string{"▓", "▒", "░"}, 100)
	add("bounce", []string{"⠁", "⠂", "⠄", "⠂"}, 120)
	add("boxBounce", []string{"▖", "▘", "▝", "▗"}, 120)
	add("boxBounce2", []string{"▌", "▀", "▐", "▄"}, 100)
	add("triangle", []string{"◢", "◣", "◤", "◥"}, 50)
	add("arc", []string{"◜", "◠", "◝", "◞", "◡", "◟"}, 100)
	add("circle", []string{"◡", "⊙", "◠"}, 120)
	add("squareCorners", []string{"◰", "◳", "◲", "◱"}, 180)
	add("circleQuarters", []string{"◴", "◷", "◶", "◵"}, 120)
	add("circleHalves", []string{"◐", "◓", "◑", "◒"}, 50)
	add("squish", []string{"╫", "╪"}, 100)
	add("toggle", []string{"⊶", "⊷"}, 250)
	add("toggle2", []string{"▫", "▪"}, 80)
	add("toggle3", []string{"□", "■"}, 120)
	add("toggle4", []string{"■", "□", "▪", "▫"}, 100)
	add("toggle5", []string{"▮", "▯"}, 100)
	add("toggle6", []string{"ဝ", "၀"}, 300)
	add("toggle7", []string{"⦾", "⦿"}, 80)
	add("toggle8", []string{"◍", "◌"}, 100)
	add("toggle9", []string{"◉", "◎"}, 100)
	add("toggle10", []string{"㊂", "㊀", "㊁"}, 100)
	add("toggle11", []string{"⧇", "⧆"}, 50)
	add("toggle12", []string{"☗", "☖"}, 120)
	add("toggle13", []string{"=", "*", "-"}, 80)
	add("arrow", []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}, 100)
	add("arrow2", []string{"⬆️ ", "↗️ ", "➡️ ", "↘️ ", "⬇️ ", "↙️ ", "⬅️ ", "↖️ "}, 80)
	add("arrow3", []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"}, 120)
	add("bouncingBar", []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"}, 80)
	add("bouncingBall", []string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"}, 80)
	add("smiley", []string{"😄 ", "😝 "}, 200)
	add("monkey", []string{"🙈 ", "🙈 ", "🙉 ", "🙊 "}, 300)
	add("hearts", []string{"💛 ", "💙 ", "💜 ", "💚 ", "❤️ "}, 100)
	add("clock", []string{"🕛 ", "🕐 ", "🕑 ", "🕒 ", "🕓 ", "🕔 ", "🕕 ", "🕖 ", "🕗 ", "🕘 ", "🕙 ", "🕚 "}, 100)
	add("earth", []string{"🌍 ", "🌎 ", "🌏 "}, 180)
	add("moon", []string{"🌑 ", "🌒 ", "🌓 ", "🌔 ", "🌕 ", "🌖 ", "🌗 ", "🌘 "}, 80)
	add("runner", []string{"🚶 ", "🏃 "}, 140)
	add("pong", []string{"▐⠂       ▌", "▐⠈       ▌", "▐ ⠂      ▌", "▐ ⠠      ▌", "▐  ⡀     ▌", "▐  ⠠     ▌", "▐   ⠂    ▌", "▐   ⠈    ▌", "▐    ⠂   ▌", "▐    ⠠   ▌", "▐     ⡀  ▌", "▐     ⠠  ▌", "▐      ⠂ ▌", "▐      ⠈ ▌", "▐       ⠂▌", "▐       ⠠▌", "▐       ⡀▌", "▐      ⠠ ▌", "▐      ⠂ ▌", "▐     ⠈  ▌", "▐     ⠂  ▌", "▐    ⠠   ▌", "▐    ⡀   ▌", "▐   ⠠    ▌", "▐   ⠂    ▌", "▐  ⠈     ▌", "▐  ⠂     ▌", "▐ ⠠      ▌", "▐ ⡀      ▌", "▐⠠       ▌"}, 80)
	add("shark", []string{"▐|\\____________▌", "▐_|\\___________▌", "▐__|\\__________▌", "▐___|\\_________▌", "▐____|\\________▌", "▐_____|\\_______▌", "▐______|\\______▌", "▐_______|\\_____▌", "▐________|\\____▌", "▐_________|\\___▌", "▐__________|\\__▌", "▐___________|\\_▌", "▐____________|\\▌", "▐____________/|▌", "▐___________/|_▌", "▐__________/|__▌", "▐_________/|___▌", "▐________/|____▌", "▐_______/|_____▌", "▐______/|______▌", "▐_____/|_______▌", "▐____/|________▌", "▐___/|_________▌", "▐__/|__________▌", "▐_/|___________▌", "▐/|____________▌"}, 120)
	add("dqpb", []string{"d", "q", "p", "b"}, 100)
	add("weather", []string{"☀️ ", "☀️ ", "☀️ ", "🌤 ", "⛅️ ", "🌥 ", "☁️ ", "🌧 ", "🌨 ", "🌧 ", "🌨 ", "🌧 ", "🌨 ", "⛈ ", "🌨 ", "🌧 ", "🌨 ", "☁️ ", "🌥 ", "⛅️ ", "🌤 ", "☀️ ", "☀️ "}, 100)
	add("christmas", []string{"🌲", "🎄"}, 400)
	add("grenade", []string{"،   ", "′   ", " ´ ", " ‾ ", "  ⸌", "  ⸊", "  |", "  ⁎", "  ⁕", " ෴ ", "  ⁓", "   ", "   ", "   "}, 80)
	add("point", []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}, 125)
	add("layer", []string{"-", "=", "≡"}, 150)
}
