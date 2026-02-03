// Code generated from ace.jar fields *.json files
// DO NOT EDIT.

package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

// just to fix compile issues with the import
var (
	_ context.Context
	_ fmt.Formatter
	_ json.Marshaler
)

const SettingUsgKey = "usg"

type SettingUsg struct {
	ID     string `json:"_id,omitempty"`
	SiteID string `json:"site_id,omitempty"`

	Hidden   bool   `json:"attr_hidden,omitempty"`
	HiddenID string `json:"attr_hidden_id,omitempty"`
	NoDelete bool   `json:"attr_no_delete,omitempty"`
	NoEdit   bool   `json:"attr_no_edit,omitempty"`

	Key string `json:"key"`

	ArpCacheBaseReachable          int                        `json:"arp_cache_base_reachable,omitempty"`                                                  // ^$|^[1-9]{1}[0-9]{0,4}$
	ArpCacheTimeout                string                     `json:"arp_cache_timeout,omitempty" validate:"omitempty,oneof=normal min-dhcp-lease custom"` // normal|min-dhcp-lease|custom
	BroadcastPing                  bool                       `json:"broadcast_ping"`
	DHCPDHostfileUpdate            bool                       `json:"dhcpd_hostfile_update"`
	DHCPDUseDNSmasq                bool                       `json:"dhcpd_use_dnsmasq"`
	DHCPRelayAgentsPackets         string                     `json:"dhcp_relay_agents_packets" validate:"omitempty,oneof=append discard forward replace"` // append|discard|forward|replace|^$
	DHCPRelayHopCount              int                        `json:"dhcp_relay_hop_count,omitempty"`                                                      // ([1-9]|[1-8][0-9]|9[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])|^$
	DHCPRelayMaxSize               int                        `json:"dhcp_relay_max_size,omitempty"`                                                       // (6[4-9]|[7-9][0-9]|[1-8][0-9]{2}|9[0-8][0-9]|99[0-9]|1[0-3][0-9]{2}|1400)|^$
	DHCPRelayPort                  int                        `json:"dhcp_relay_port,omitempty"`                                                           // [1-9][0-9]{0,3}|[1-5][0-9]{4}|[6][0-4][0-9]{3}|[6][5][0-4][0-9]{2}|[6][5][5][0-2][0-9]|[6][5][5][3][0-5]|^$
	DHCPRelayServer1               string                     `json:"dhcp_relay_server_1" validate:"omitempty,ipv4"`                                       // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	DHCPRelayServer2               string                     `json:"dhcp_relay_server_2" validate:"omitempty,ipv4"`                                       // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	DHCPRelayServer3               string                     `json:"dhcp_relay_server_3" validate:"omitempty,ipv4"`                                       // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	DHCPRelayServer4               string                     `json:"dhcp_relay_server_4" validate:"omitempty,ipv4"`                                       // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	DHCPRelayServer5               string                     `json:"dhcp_relay_server_5" validate:"omitempty,ipv4"`                                       // ^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$
	DNSVerification                SettingUsgDNSVerification  `json:"dns_verification,omitempty"`
	DNSmasqAllServers              bool                       `json:"dnsmasq_all_servers"`
	EchoServer                     string                     `json:"echo_server,omitempty"` // [^\"\' ]{1,255}
	FtpModule                      bool                       `json:"ftp_module"`
	GeoIPFilteringBlock            string                     `json:"geo_ip_filtering_block,omitempty" validate:"omitempty,oneof=block allow"` // block|allow
	GeoIPFilteringCountries        string                     `json:"geo_ip_filtering_countries"`                                              // ^([A-Z]{2})?(,[A-Z]{2}){0,149}$
	GeoIPFilteringEnabled          bool                       `json:"geo_ip_filtering_enabled"`
	GeoIPFilteringTrafficDirection string                     `json:"geo_ip_filtering_traffic_direction,omitempty" validate:"omitempty,oneof=both ingress egress"` // ^(both|ingress|egress)$
	GreModule                      bool                       `json:"gre_module"`
	H323Module                     bool                       `json:"h323_module"`
	ICMPTimeout                    int                        `json:"icmp_timeout,omitempty"`
	LldpEnableAll                  bool                       `json:"lldp_enable_all"`
	MdnsEnabled                    bool                       `json:"mdns_enabled"`
	MssClamp                       string                     `json:"mss_clamp,omitempty" validate:"omitempty,oneof=auto custom disabled"` // auto|custom|disabled
	MssClampMss                    int                        `json:"mss_clamp_mss,omitempty"`                                             // [1-9][0-9]{2,3}
	OffloadAccounting              bool                       `json:"offload_accounting"`
	OffloadL2Blocking              bool                       `json:"offload_l2_blocking"`
	OffloadSch                     bool                       `json:"offload_sch"`
	OtherTimeout                   int                        `json:"other_timeout,omitempty"`
	PptpModule                     bool                       `json:"pptp_module"`
	ReceiveRedirects               bool                       `json:"receive_redirects"`
	SendRedirects                  bool                       `json:"send_redirects"`
	SipModule                      bool                       `json:"sip_module"`
	SynCookies                     bool                       `json:"syn_cookies"`
	TCPCloseTimeout                int                        `json:"tcp_close_timeout,omitempty"`
	TCPCloseWaitTimeout            int                        `json:"tcp_close_wait_timeout,omitempty"`
	TCPEstablishedTimeout          int                        `json:"tcp_established_timeout,omitempty"`
	TCPFinWaitTimeout              int                        `json:"tcp_fin_wait_timeout,omitempty"`
	TCPLastAckTimeout              int                        `json:"tcp_last_ack_timeout,omitempty"`
	TCPSynRecvTimeout              int                        `json:"tcp_syn_recv_timeout,omitempty"`
	TCPSynSentTimeout              int                        `json:"tcp_syn_sent_timeout,omitempty"`
	TCPTimeWaitTimeout             int                        `json:"tcp_time_wait_timeout,omitempty"`
	TFTPModule                     bool                       `json:"tftp_module"`
	TimeoutSettingPreference       string                     `json:"timeout_setting_preference,omitempty" validate:"omitempty,oneof=auto manual"` // auto|manual
	UDPOtherTimeout                int                        `json:"udp_other_timeout,omitempty"`
	UDPStreamTimeout               int                        `json:"udp_stream_timeout,omitempty"`
	UnbindWANMonitors              bool                       `json:"unbind_wan_monitors"`
	UpnpEnabled                    bool                       `json:"upnp_enabled"`
	UpnpNATPmpEnabled              bool                       `json:"upnp_nat_pmp_enabled"`
	UpnpSecureMode                 bool                       `json:"upnp_secure_mode"`
	UpnpWANInterface               string                     `json:"upnp_wan_interface,omitempty" validate:"omitempty,oneof=WAN WAN2"` // WAN|WAN2
	ExtraFields                    map[string]json.RawMessage `json:"-"`
}

func (dst *SettingUsg) UnmarshalJSON(b []byte) error {
	type Alias SettingUsg
	aux := &struct {
		ArpCacheBaseReachable emptyStringInt `json:"arp_cache_base_reachable"`
		DHCPRelayHopCount     emptyStringInt `json:"dhcp_relay_hop_count"`
		DHCPRelayMaxSize      emptyStringInt `json:"dhcp_relay_max_size"`
		DHCPRelayPort         emptyStringInt `json:"dhcp_relay_port"`
		ICMPTimeout           emptyStringInt `json:"icmp_timeout"`
		MssClampMss           emptyStringInt `json:"mss_clamp_mss"`
		OtherTimeout          emptyStringInt `json:"other_timeout"`
		TCPCloseTimeout       emptyStringInt `json:"tcp_close_timeout"`
		TCPCloseWaitTimeout   emptyStringInt `json:"tcp_close_wait_timeout"`
		TCPEstablishedTimeout emptyStringInt `json:"tcp_established_timeout"`
		TCPFinWaitTimeout     emptyStringInt `json:"tcp_fin_wait_timeout"`
		TCPLastAckTimeout     emptyStringInt `json:"tcp_last_ack_timeout"`
		TCPSynRecvTimeout     emptyStringInt `json:"tcp_syn_recv_timeout"`
		TCPSynSentTimeout     emptyStringInt `json:"tcp_syn_sent_timeout"`
		TCPTimeWaitTimeout    emptyStringInt `json:"tcp_time_wait_timeout"`
		UDPOtherTimeout       emptyStringInt `json:"udp_other_timeout"`
		UDPStreamTimeout      emptyStringInt `json:"udp_stream_timeout"`

		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}
	dst.ArpCacheBaseReachable = int(aux.ArpCacheBaseReachable)
	dst.DHCPRelayHopCount = int(aux.DHCPRelayHopCount)
	dst.DHCPRelayMaxSize = int(aux.DHCPRelayMaxSize)
	dst.DHCPRelayPort = int(aux.DHCPRelayPort)
	dst.ICMPTimeout = int(aux.ICMPTimeout)
	dst.MssClampMss = int(aux.MssClampMss)
	dst.OtherTimeout = int(aux.OtherTimeout)
	dst.TCPCloseTimeout = int(aux.TCPCloseTimeout)
	dst.TCPCloseWaitTimeout = int(aux.TCPCloseWaitTimeout)
	dst.TCPEstablishedTimeout = int(aux.TCPEstablishedTimeout)
	dst.TCPFinWaitTimeout = int(aux.TCPFinWaitTimeout)
	dst.TCPLastAckTimeout = int(aux.TCPLastAckTimeout)
	dst.TCPSynRecvTimeout = int(aux.TCPSynRecvTimeout)
	dst.TCPSynSentTimeout = int(aux.TCPSynSentTimeout)
	dst.TCPTimeWaitTimeout = int(aux.TCPTimeWaitTimeout)
	dst.UDPOtherTimeout = int(aux.UDPOtherTimeout)
	dst.UDPStreamTimeout = int(aux.UDPStreamTimeout)

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"_id":                                {},
			"site_id":                            {},
			"attr_hidden":                        {},
			"attr_hidden_id":                     {},
			"attr_no_delete":                     {},
			"attr_no_edit":                       {},
			"key":                                {},
			"arp_cache_base_reachable":           {},
			"arp_cache_timeout":                  {},
			"broadcast_ping":                     {},
			"dhcpd_hostfile_update":              {},
			"dhcpd_use_dnsmasq":                  {},
			"dhcp_relay_agents_packets":          {},
			"dhcp_relay_hop_count":               {},
			"dhcp_relay_max_size":                {},
			"dhcp_relay_port":                    {},
			"dhcp_relay_server_1":                {},
			"dhcp_relay_server_2":                {},
			"dhcp_relay_server_3":                {},
			"dhcp_relay_server_4":                {},
			"dhcp_relay_server_5":                {},
			"dns_verification":                   {},
			"dnsmasq_all_servers":                {},
			"echo_server":                        {},
			"ftp_module":                         {},
			"geo_ip_filtering_block":             {},
			"geo_ip_filtering_countries":         {},
			"geo_ip_filtering_enabled":           {},
			"geo_ip_filtering_traffic_direction": {},
			"gre_module":                         {},
			"h323_module":                        {},
			"icmp_timeout":                       {},
			"lldp_enable_all":                    {},
			"mdns_enabled":                       {},
			"mss_clamp":                          {},
			"mss_clamp_mss":                      {},
			"offload_accounting":                 {},
			"offload_l2_blocking":                {},
			"offload_sch":                        {},
			"other_timeout":                      {},
			"pptp_module":                        {},
			"receive_redirects":                  {},
			"send_redirects":                     {},
			"sip_module":                         {},
			"syn_cookies":                        {},
			"tcp_close_timeout":                  {},
			"tcp_close_wait_timeout":             {},
			"tcp_established_timeout":            {},
			"tcp_fin_wait_timeout":               {},
			"tcp_last_ack_timeout":               {},
			"tcp_syn_recv_timeout":               {},
			"tcp_syn_sent_timeout":               {},
			"tcp_time_wait_timeout":              {},
			"tftp_module":                        {},
			"timeout_setting_preference":         {},
			"udp_other_timeout":                  {},
			"udp_stream_timeout":                 {},
			"unbind_wan_monitors":                {},
			"upnp_enabled":                       {},
			"upnp_nat_pmp_enabled":               {},
			"upnp_secure_mode":                   {},
			"upnp_wan_interface":                 {},
		}
		for k, v := range raw {
			if _, ok := known[k]; !ok {
				if dst.ExtraFields == nil {
					dst.ExtraFields = make(map[string]json.RawMessage)
				}
				dst.ExtraFields[k] = v
			}
		}
	}

	return nil
}

func (src SettingUsg) MarshalJSON() ([]byte, error) {
	type Alias SettingUsg
	b, err := json.Marshal(Alias(src))
	if err != nil {
		return nil, err
	}
	if len(src.ExtraFields) == 0 {
		return b, nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	extra, err := json.Marshal(src.ExtraFields)
	if err != nil {
		return nil, err
	}
	m["_additional_properties"] = extra
	return json.Marshal(m)
}

type SettingUsgDNSVerification struct {
	Domain             string                     `json:"domain,omitempty"`
	PrimaryDNSServer   string                     `json:"primary_dns_server,omitempty"`
	SecondaryDNSServer string                     `json:"secondary_dns_server,omitempty"`
	SettingPreference  string                     `json:"setting_preference,omitempty" validate:"omitempty,oneof=auto manual"` // auto|manual
	ExtraFields        map[string]json.RawMessage `json:"-"`
}

func (dst *SettingUsgDNSVerification) UnmarshalJSON(b []byte) error {
	type Alias SettingUsgDNSVerification
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dst),
	}

	err := json.Unmarshal(b, &aux)
	if err != nil {
		return fmt.Errorf("unable to unmarshal alias: %w", err)
	}

	// Capture extra fields not in the struct
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err == nil {
		known := map[string]struct{}{
			"domain":               {},
			"primary_dns_server":   {},
			"secondary_dns_server": {},
			"setting_preference":   {},
		}
		for k, v := range raw {
			if _, ok := known[k]; !ok {
				if dst.ExtraFields == nil {
					dst.ExtraFields = make(map[string]json.RawMessage)
				}
				dst.ExtraFields[k] = v
			}
		}
	}

	return nil
}

func (src SettingUsgDNSVerification) MarshalJSON() ([]byte, error) {
	type Alias SettingUsgDNSVerification
	b, err := json.Marshal(Alias(src))
	if err != nil {
		return nil, err
	}
	if len(src.ExtraFields) == 0 {
		return b, nil
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	extra, err := json.Marshal(src.ExtraFields)
	if err != nil {
		return nil, err
	}
	m["_additional_properties"] = extra
	return json.Marshal(m)
}

// GetSettingUsg Experimental! This function is not yet stable and may change in the future.
func (c *client) GetSettingUsg(ctx context.Context, site string) (*SettingUsg, error) {
	s, f, err := c.GetSetting(ctx, site, SettingUsgKey)
	if err != nil {
		return nil, err
	}
	if s.Key != SettingUsgKey {
		return nil, fmt.Errorf("unexpected setting key received. Requested: %q, received: %q", SettingUsgKey, s.Key)
	}
	return f.(*SettingUsg), nil
}

// UpdateSettingUsg Experimental! This function is not yet stable and may change in the future.
func (c *client) UpdateSettingUsg(ctx context.Context, site string, s *SettingUsg) (*SettingUsg, error) {
	s.Key = SettingUsgKey
	result, err := c.SetSetting(ctx, site, SettingUsgKey, s)
	if err != nil {
		return nil, err
	}
	return result.(*SettingUsg), nil
}
