#include <git2.h>
#include <stdio.h>
#include <string.h>

static void commit_parsing(git_repository *repo);
static void check_error(int error_code, const char *action);

int main(int argc, char **argv)
{
	git_repository *repo;
	int error;
	git_oid oid;
	char *repo_path;

	/**
	 * Initialize the library, this will set up any global state which libgit2 needs
	 * including threading and crypto
	 */
	git_libgit2_init();

	/**
	 * ### Opening the Repository
	 *
	 * There are a couple of methods for opening a repository, this being the
	 * simplest.  There are also [methods][me] for specifying the index file
	 * and work tree locations, here we assume they are in the normal places.
	 *
	 * (Try running this program against tests/resources/testrepo.git.)
	 *
	 * [me]: http://libgit2.github.com/libgit2/#HEAD/group/repository
	 */
	repo_path = "./";

	error = git_repository_open(&repo, repo_path);
	check_error(error, "opening repository");

	// oid_parsing(&oid);
	// object_database(repo, &oid);
	// commit_writing(repo);
	commit_parsing(repo);
	// tag_parsing(repo);
	// tree_parsing(repo);
	// blob_parsing(repo);
	// revwalking(repo);
	// index_walking(repo);
	// reference_listing(repo);
	// config_files(repo_path, repo);

	/**
	 * Finally, when you're done with the repository, you can free it as well.
	 */
	git_repository_free(repo);

	return 0;
}

/**
 * #### Commit Parsing
 *
 * [Parsing commit objects][pco] is simple and gives you access to all the
 * data in the commit - the author (name, email, datetime), committer
 * (same), tree, message, encoding and parent(s).
 *
 * [pco]: http://libgit2.github.com/libgit2/#HEAD/group/commit
 */
static void commit_parsing(git_repository *repo)
{
	const git_signature *author, *cmtter;
	git_commit *commit, *parent;
	git_oid oid;
	char oid_hex[GIT_OID_HEXSZ + 1];
	const char *message;
	unsigned int parents, p;
	int error;
	time_t time;

	printf("\n*Commit Parsing*\n");

	git_oid_fromstr(&oid, "8496071c1b46c854b31185ea97743be6a8774479");

	error = git_commit_lookup(&commit, repo, &oid);
	check_error(error, "looking up commit");

	/**
	 * Each of the properties of the commit object are accessible via methods,
	 * including commonly needed variations, such as `git_commit_time` which
	 * returns the author time and `git_commit_message` which gives you the
	 * commit message (as a NUL-terminated string).
	 */
	message = git_commit_message(commit);
	author = git_commit_author(commit);
	cmtter = git_commit_committer(commit);
	time = git_commit_time(commit);

	/**
	 * The author and committer methods return [git_signature] structures,
	 * which give you name, email and `when`, which is a `git_time` structure,
	 * giving you a timestamp and timezone offset.
	 */
	printf("Author: %s (%s)\nCommitter: %s (%s)\nDate: %s\nMessage: %s\n",
		   author->name, author->email,
		   cmtter->name, cmtter->email,
		   ctime(&time), message);

	/**
	 * Commits can have zero or more parents. The first (root) commit will
	 * have no parents, most commits will have one (i.e. the commit it was
	 * based on) and merge commits will have two or more.  Commits can
	 * technically have any number, though it's rare to have more than two.
	 */
	parents = git_commit_parentcount(commit);
	for (p = 0; p < parents; p++)
	{
		memset(oid_hex, 0, sizeof(oid_hex));

		git_commit_parent(&parent, commit, p);
		git_oid_fmt(oid_hex, git_commit_id(parent));
		printf("Parent: %s\n", oid_hex);
		git_commit_free(parent);
	}

	git_commit_free(commit);
}

static void check_error(int error_code, const char *action)
{
	const git_error *error = git_error_last();
	if (!error_code)
		return;

	printf("Error %d %s - %s\n", error_code, action,
		   (error && error->message) ? error->message : "???");

	exit(1);
}