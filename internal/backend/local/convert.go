/*
Sniperkit-Bot
- Status: analyzed
*/

package local

func toEnv(env map[string]string) []string {
	var envs []string
	for k, v := range env {
		envs = append(envs, k+"="+v)
	}
	return envs
}
