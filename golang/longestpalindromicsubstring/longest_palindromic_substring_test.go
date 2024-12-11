package longestpalindromicsubstring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := ""
	expected := ""
	result := longestPalindrome(s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}

	s = "a"
	expected = "a"
	result = longestPalindrome(s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}

	s = "ac"
	expected = "a"
	result = longestPalindrome(s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}

	s = "babad"
	expected = "bab"
	result = longestPalindrome(s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}

	s = "cbbd"
	expected = "bb"
	result = longestPalindrome(s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}

	s = "uzhynqvopdbnkvuxizirzjsslptlhmvyfqhqvqffmqldkrrdwapbdcxqbchvxqixmvbbtalrgzvkobyxlkonkfknabjwzeoazankqzuhexhcwkbvwtioubrcujqqeqoedhploqklhgeilwwnndbnzeadzefkcvaxdhgnmocadvvzjocoweyoidhleuwhmavpdiizbwkukpstyorlwwyiwwyyyzqqgipzzlxsgcdjscdfvrrrqmllbdjkkuisxeqaprfblvjuixajhucmcvffmevaztvadrujbthlnsdrxzvbldwxbazxmilpkbccigkihcgjbtpvignmrgzdqnufaacxtihfgwrllrwgfhitxcaafunqdzgrmngivptbjgchikgiccbkplimxzabxwdlbvzxrdsnlhtbjurdavtzavemffvcmcuhjaxiujvlbfrpaqexsiukkjdbllmqrrrvfdcsjdcgsxlzzpigqqzyyywwiywwlroytspkukwbziidpvamhwuelhdioyewocojzvvdacomnghdxavckfezdaeznbdnnwwlieghlkqolphdeoqeqqjucrbuoitwvbkwchxehuzqknazaoezwjbankfknoklxybokvzgrlatbbvmxiqxvhcbqxcdbpawdrrkdlqmffqvqhqfyvmhltplssjzrizixuvknbdpovqnyhzu"
	expected = "uzhynqvopdbnkvuxizirzjsslptlhmvyfqhqvqffmqldkrrdwapbdcxqbchvxqixmvbbtalrgzvkobyxlkonkfknabjwzeoazankqzuhexhcwkbvwtioubrcujqqeqoedhploqklhgeilwwnndbnzeadzefkcvaxdhgnmocadvvzjocoweyoidhleuwhmavpdiizbwkukpstyorlwwyiwwyyyzqqgipzzlxsgcdjscdfvrrrqmllbdjkkuisxeqaprfblvjuixajhucmcvffmevaztvadrujbthlnsdrxzvbldwxbazxmilpkbccigkihcgjbtpvignmrgzdqnufaacxtihfgwrllrwgfhitxcaafunqdzgrmngivptbjgchikgiccbkplimxzabxwdlbvzxrdsnlhtbjurdavtzavemffvcmcuhjaxiujvlbfrpaqexsiukkjdbllmqrrrvfdcsjdcgsxlzzpigqqzyyywwiywwlroytspkukwbziidpvamhwuelhdioyewocojzvvdacomnghdxavckfezdaeznbdnnwwlieghlkqolphdeoqeqqjucrbuoitwvbkwchxehuzqknazaoezwjbankfknoklxybokvzgrlatbbvmxiqxvhcbqxcdbpawdrrkdlqmffqvqhqfyvmhltplssjzrizixuvknbdpovqnyhzu"
	result = longestPalindrome(s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}

	s = "busislnescsicxpvvysuqgcudefrfjbwwjcchtgqyajdfwvkypfwshnihjdztgmyuuljxgvhdiwphrweyfkbnjgerkmifbirubhseuhrugwrabnjafnbdfjnufdstjbkuwtnpflffaqmjbhssjlnqftgjiglvvequhapasarlkcvbmkwnkuvwktbgfoaxteprobdwswcdyddyvrehvmxrrjiiidatidlpihkbmmruysmhhsncmfdanafdrfpdtfgkglcqpwrrtvacuicohspkounojuziittugpqjyhhkwfnflozbispehrtrnizowrlzcuollagxwtznjwzcumvedjwokueuqktvvouwnsmpxqvvpuwprezrbobrpnwaccwljchdguubjulyilzvmandjjleitweybqkjttschrjjlebnmponvlktzzcdtuybugggcqffkcffpamauvxfbonjrobgpvlyzveiwemmtdvbjciaytvesnocnjrwodtcokgcuoiicxapmrzpkfphjniuvzjrhbnqndfshoduejyktebgdabidxlkstepuwvtrtgbxaeheylicvhrxddijshcvdadxzsccmainyfpfdhqdanfccqkzlmhsfilvoybqojlvbcixjzqpbngdvesuokbxhkomsiqfyukvspqthlzxdnlwthrgaxhtpjzhrugqbfokrdcyurivmzgtynoqfjbafboselxnfupnpqlryvlcxeksirvufepfwczosrrjpudbwqxwldgjyfjhzlzcojxyqjyxxiqvfhjdwtgoqbyeocffnyxhyyiqspnvrpxmrtcnviukrjvpavervvztoxajriuvxqveqsrttjqepvvahywuzwtmgyrzduxfqspeipimyoxmkadrvrdyefekjxcmsmzmtbugyckcbjsrymszftjyllfmoeoylzeahnrxlxpnlvlvzltwnmldi"
	expected = "rbobr"
	result = longestPalindrome(s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %s, expected %s", s, result, expected)
	}
}

func TestIsPalindrome(t *testing.T) {
	s := ""
	expected := true
	result := isPalindrome(&s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %v, expected %v", s, result, expected)
	}

	s = "anavolimilovana"
	expected = true
	result = isPalindrome(&s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %v, expected %v", s, result, expected)
	}

	s = "ananevolimilovana"
	expected = false
	result = isPalindrome(&s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %v, expected %v", s, result, expected)
	}

	s = "wontloversrevoltnow"
	expected = true
	result = isPalindrome(&s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %v, expected %v", s, result, expected)
	}

	s = "loverswontrevoltnow"
	expected = false
	result = isPalindrome(&s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %v, expected %v", s, result, expected)
	}

	s = "uzhynqvopdbnkvuxizirzjsslptlhmvyfqhqvqffmqldkrrdwapbdcxqbchvxqixmvbbtalrgzvkobyxlkonkfknabjwzeoazankqzuhexhcwkbvwtioubrcujqqeqoedhploqklhgeilwwnndbnzeadzefkcvaxdhgnmocadvvzjocoweyoidhleuwhmavpdiizbwkukpstyorlwwyiwwyyyzqqgipzzlxsgcdjscdfvrrrqmllbdjkkuisxeqaprfblvjuixajhucmcvffmevaztvadrujbthlnsdrxzvbldwxbazxmilpkbccigkihcgjbtpvignmrgzdqnufaacxtihfgwrllrwgfhitxcaafunqdzgrmngivptbjgchikgiccbkplimxzabxwdlbvzxrdsnlhtbjurdavtzavemffvcmcuhjaxiujvlbfrpaqexsiukkjdbllmqrrrvfdcsjdcgsxlzzpigqqzyyywwiywwlroytspkukwbziidpvamhwuelhdioyewocojzvvdacomnghdxavckfezdaeznbdnnwwlieghlkqolphdeoqeqqjucrbuoitwvbkwchxehuzqknazaoezwjbankfknoklxybokvzgrlatbbvmxiqxvhcbqxcdbpawdrrkdlqmffqvqhqfyvmhltplssjzrizixuvknbdpovqnyhzu"
	expected = true
	result = isPalindrome(&s)

	if result != expected {
		t.Errorf("lengthOfLongestSubstring(%s) returned %v, expected %v", s, result, expected)
	}
}

