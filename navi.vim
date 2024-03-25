" Integrates Navi into Vim interface.
" Ctrl+K will prompt the cursor line.

fun! Navi(s)
	let current_line = line('.')
	let s = escape(a:s, "\"")
	let s = escape(a:s, "\'")
	let s = escape(a:s, "\`")
	let s = escape(a:s, "\\")
	let out = system("_ \'" . s . "\' | sed \'s/\x1B\[[0-9;]*[mK]//g\'")
	let out = substitute(out, '>> ', '', 'g')
	let out = substitute(out, '>> ', '', 'g')
	let out = trim(out)
	let lines = split(out, '\n')
	call append(current_line, '')
	let current_line += 1
	for line in lines
		call append(current_line, line)
		let current_line += 1
	endfor
	call append(current_line, '')
endfunction

nnoremap <C-K> :call Navi(getline('.'))<CR>

