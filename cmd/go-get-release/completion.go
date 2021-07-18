package main

import "fmt"

var bashCompletion = `
__go_get_release_bash_autocomplete() {
  if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    if [[ "$cur" == "-"* ]]; then
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    else
      opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    fi
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  fi
}

complete -o bashdefault -o default -o nospace -F __go_get_release_bash_autocomplete go-get-release
`

func completion(shell string) error {
	if shell != "bash" {
		return fmt.Errorf("supported shell is bash only")
	}
	fmt.Println(bashCompletion)
	return nil
}
