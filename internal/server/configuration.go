package server

import (
	"github.com/gavrylenkoIvan/hopper/internal/hopper"
)

// Handle Configuration state
// https://wiki.vg/Protocol#Configuration
func (h *Hopper) handleConfiguration(conn *hopper.Conn) error {
	// buf := bytes.NewBuffer(nil)
	// _, err := types.VarInt(0x05).WriteTo(buf)
	// if err != nil {
	// 	return err
	// }

	// _, err = buf.Write(registry.Codec)
	// if err != nil {
	// 	return err
	// }

	// _, err = conn.WritePacket(buf.Bytes())
	// if err != nil {
	// 	return err
	// }

	// finishConfiguration, err := packet.Marshal(
	// 	types.VarInt(0x02),
	// )
	// if err != nil {
	// 	return err
	// }

	// _, err = conn.WritePacket(finishConfiguration)

	return nil
}
