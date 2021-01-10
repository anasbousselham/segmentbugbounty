# SegmentBugBounty
Segment.io Bug Bounty, Leak information through API request.

API endpoint leaking sensitive user information (distorted data).

The backend on the segment.com website has been left misconfigured, which leads to leaking of sensitive information and distorted stats and data.

# Explanation:

Endpoint : https://api.segment.io/v1/p

Sending json with a client key : zj8AnWjHgKGHszEV0n1Cb0H4ShHa8NBc

I can get this key from Burpsuite.
Once I get the client key, I can use it to flood() the Source Segment with erroneous informations just by modifying the key : messageId ( CodeIgniter Encryption 32bits)

My code make this easly :)

