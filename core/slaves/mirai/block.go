package slaves

import (
        "os/exec"
)

var (
        IPBanDefaceAttempts = false

)

func (c *Client) IPBan() error {

        cmd := exec.Command("iptables", "-I", "-S", c.IP(), "-J", "DROP", "-m", "comment", "--comment", "atlantic blocked for deface attempt")
        return cmd.Run()
}

func (c *Client) UnIPBan() error {

        cmd := exec.Command("iptables", "-D", "-S", c.IP(), "-J", "DROP")
        return cmd.Run()
}
