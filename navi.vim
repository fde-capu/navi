" Integrates Navi into Vim interface.
" Ctrl+K will prompt the cursor line.
" In Visual mode, will prompt the select portion.

let g:NaviBin = '_'

function! EscapeSpecialChars(str)
	let escapedStr = substitute(a:str, "'", "''", 'g')
	let escapedStr = substitute(escapedStr, '"', '\\"', 'g')
	let escapedStr = substitute(escapedStr, "`", "\\`", 'g')
	let escapedStr = substitute(escapedStr, [^a-zA-Z0-9], '', g)
	return escapedStr
endfunction

fun! Navi(s)
	let current_line = line('.')
	echo "Original string: " . a:s
	let s = EscapeSpecialChars(a:s)
	echo "Escaped string: " . s
	let out = system(g:NaviBin . " \'" . s . "\'")
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

function! GetSelectionOrCurrentLine()
	if visualmode() !=# ''
		echo 'Visual mode'
		let start_pos = getpos("'<") " Vim Script
		let end_pos = getpos("'>")
		let start_line = start_pos[1]
		let end_line = end_pos[1]
		let lines = getline(start_line, end_line)
		call Navi(join(lines, "\n"))
	else
		let line = getline('.')
		call Navi(line)
	endif
endfunction

nnoremap <C-K> :call GetSelectionOrCurrentLine()<CR>
vnoremap <C-K> :call GetSelectionOrCurrentLine()<CR>
