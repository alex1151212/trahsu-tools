import instaloader
import sys

username = sys.argv[1]


if username:
    L = instaloader.Instaloader()
    profile = instaloader.Profile.from_username(L.context, username)
    numPhoto = 1
    for post in profile.get_posts():
        numPhoto += post.mediacount
        for node in post.get_sidecar_nodes():
            print(node.display_url)
