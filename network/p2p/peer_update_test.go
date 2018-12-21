/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */
package p2p

import (
	"reflect"
	"testing"

	libnetwork "github.com/creativesoftwarefdn/weaviate/network"
)

func TestPeerUpdateWithNewPeers(t *testing.T) {
	oldPeers := []libnetwork.Peer{}
	newPeers := []libnetwork.Peer{{
		Name: "best-weaviate",
		Id:   "uuid",
		URI:  "does-not-matter",
	}}

	subject := network{
		peers: oldPeers,
	}

	callbackCalled := false
	callbackCalledWith := []libnetwork.Peer{}
	callbackSpy := func(peers libnetwork.Peers) {
		callbackCalled = true
		callbackCalledWith = peers
	}

	subject.RegisterUpdatePeerCallback(callbackSpy)
	subject.UpdatePeers(newPeers)

	if callbackCalled != true {
		t.Error("expect PeerUpdateCallback to be called, but was never called")
	}

	if !reflect.DeepEqual(callbackCalledWith, newPeers) {
		t.Errorf("expect PeerUpdateCallback to be called with new peers, but was called with %#v",
			callbackCalledWith)
	}
}

func TestPeerUpdateWithoutAnyChange(t *testing.T) {
	peers := []libnetwork.Peer{{
		Name: "best-weaviate",
		Id:   "uuid",
		URI:  "does-not-matter",
	}}

	subject := network{
		peers: peers,
	}

	callbackCalled := false
	callbackSpy := func(peers libnetwork.Peers) {
		callbackCalled = true
	}

	subject.RegisterUpdatePeerCallback(callbackSpy)
	subject.UpdatePeers(peers)

	if callbackCalled != false {
		t.Error("expect PeerUpdateCallback not to be called, but it was called")
	}
}