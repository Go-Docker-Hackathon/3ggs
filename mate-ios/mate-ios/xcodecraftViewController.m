//
//  xcodecraftViewController.m
//  mate-ios
//
//  Created by yangwm on 6/6/15.
//  Copyright (c) 2015 xcodecraft. All rights reserved.
//

#import "xcodecraftViewController.h"

@interface xcodecraftViewController ()
- (IBAction)login:(id)sender;

@end

@implementation xcodecraftViewController

- (void)viewDidLoad
{
    [super viewDidLoad];
	// Do any additional setup after loading the view, typically from a nib.
}

- (void)didReceiveMemoryWarning
{
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (IBAction)login:(id)sender {
    
    [[[UIAlertView alloc]initWithTitle:@"title" message:@"登录成功" delegate:nil cancelButtonTitle:@"取消" otherButtonTitles:nil, nil]show];
}
@end
